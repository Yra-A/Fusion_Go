package jwt

import (
	"context"
	"fmt"
	"github.com/Yra-A/Fusion_Go/cmd/api/biz/handler"
	"github.com/Yra-A/Fusion_Go/cmd/api/biz/model/api"
	"github.com/Yra-A/Fusion_Go/cmd/api/rpc"
	"github.com/Yra-A/Fusion_Go/kitex_gen/user"
	"github.com/Yra-A/Fusion_Go/pkg/constants"
	"github.com/Yra-A/Fusion_Go/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/jwt"
	"net/http"
	"time"
)

var JwtMiddleware *jwt.HertzJWTMiddleware

func InitJwt() {
	var err error
	JwtMiddleware, err = jwt.New(&jwt.HertzJWTMiddleware{
		Key:           []byte(constants.SecretKey),
		TimeFunc:      time.Now,
		Timeout:       7 * 24 * time.Hour,      // access token 过期时间 7 天
		MaxRefresh:    30 * 24 * time.Hour,     // refresh 过期时间为 30 天
		TokenLookup:   "header: Authorization", // 设置 token 的获取源
		TokenHeadName: "Bearer",                // 设置从 header 中获取 token 时的前缀
		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
			hlog.CtxInfof(ctx, "Login success ，token is issued clientIP: "+c.ClientIP())
			handler.SendResponse(c, api.UserLoginResponse{
				StatusCode: errno.Success.ErrCode,
				StatusMsg:  errno.Success.ErrMsg,
				Token:      token,
			})
		},
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			fmt.Println("Authenticator 成功")
			var err error
			var req api.UserLoginRequest
			if err = c.BindAndValidate(&req); err != nil {
				return nil, jwt.ErrMissingLoginValues
			}
			if len(req.Username) == 0 || len(req.Password) == 0 {
				return nil, jwt.ErrMissingLoginValues
			}
			kresp, err := rpc.UserLogin(context.Background(), &user.UserLoginRequest{
				Username: req.Username,
				Password: req.Password,
			})
			if err != nil {
				handler.BadResponse(c, err)
				return nil, err
			}
			// 检查用户登录是否成功
			if kresp.StatusCode != errno.Success.ErrCode {
				return nil, jwt.ErrFailedAuthentication
			}
			return kresp.UserId, nil // 将 UserId 存入 token 的负载部分
		},
		Authorizator: func(data interface{}, ctx context.Context, c *app.RequestContext) bool {
			//验证已登录成功的用户 data(user_id) 的权限
			path := c.FullPath()
			var userId int32
			if floatId, ok := data.(float64); ok {
				userId = int32(floatId)
			}
			if path == "/fusion/user/info" {
				var req api.UserInfoRequest
				if err = c.BindAndValidate(&req); err != nil {
					return false
				}
				if userId != req.UserID {
					return false
				}
			} else if path == "/fusion/user/profile/:user_id" {
				var req api.UserProfileInfoRequest
				if err = c.BindAndValidate(&req); err != nil {
					return false
				}
				if userId != req.UserID {
					return false
				}
			} else if path == "/fusion/user/profile/upload" {
				var req api.UserProfileUploadRequest
				if err = c.BindAndValidate(&req); err != nil {
					return false
				}
				if userId != req.UserID || userId != req.UserProfileInfo.UserInfo.UserID {
					return false
				}
			} else if path == "/fusion/team/create" {
				var req api.TeamCreateRequest
				if err = c.BindAndValidate(&req); err != nil {
					return false
				}
				if userId != req.UserID {
					return false
				}
			} else if path == "/fusion/team/application/submit" {
				var req api.TeamApplicationSubmitRequest
				if err = c.BindAndValidate(&req); err != nil {
					return false
				}
				if userId != req.MemberInfo.UserID {
					return false
				}
			} else if path == "/fusion/team/manage/list" {
				var req api.TeamManageListRequest
				if err = c.BindAndValidate(&req); err != nil {
					return false
				}
				if req.UserID != userId {
					return false
				}
			} else if path == "/fusion/team/manage/action" {
				var req api.TeamManageActionRequest
				if err = c.BindAndValidate(&req); err != nil {
					return false
				}
				if req.UserID != userId {
					return false
				}
			} else if path == "/fusion/favorite/contest/action" {
				var req api.ContestFavoriteActionRequest
				if err = c.BindAndValidate(&req); err != nil {
					return false
				}
				if req.UserID != userId {
					return false
				}
			} else if path == "/fusion/favorite/contest/list" {
				var req api.ContestFavoriteListRequest
				if err = c.BindAndValidate(&req); err != nil {
					return false
				}
				if req.UserID != userId {
					return false
				}
			}
			return true
		},
		IdentityKey: constants.IdentityKey,

		PayloadFunc: func(data interface{}) jwt.MapClaims {
			// 将 UserId 存入 token 的负载部分
			if v, ok := data.(int32); ok {
				return jwt.MapClaims{
					constants.IdentityKey: v,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			// 登陆成功后的每次请求会先执行中间件，先检查过期时间，然后会将 UserId 从 token 的负载部分取出，并存入上下文
			claims := jwt.ExtractClaims(ctx, c)
			return claims[constants.IdentityKey]
		},
		HTTPStatusMessageFunc: func(e error, ctx context.Context, c *app.RequestContext) string {
			hlog.CtxErrorf(ctx, "jwt biz err = %+v", e.Error())
			return e.Error()
		},
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			path := c.FullPath()
			if path == "/fusion/user/login" {
				c.JSON(http.StatusUnauthorized, handler.Response{
					StatusCode: errno.InvalidCredentialsErr.ErrCode,
					StatusMsg:  errno.InvalidCredentialsErr.ErrMsg,
				})
			} else {
				c.JSON(http.StatusUnauthorized, handler.Response{
					StatusCode: errno.AuthorizationFailedErr.ErrCode,
					StatusMsg:  errno.AuthorizationFailedErr.ErrMsg,
				})
			}
		},
		//ParseOptions: []jwtv4.ParserOption{
		//	jwtv4.WithValidMethods([]string{"HS256"}),
		//	//jwtv4.WithJSONNumber(),
		//},
	})
	if err != nil {
		panic(err)
	}
}
