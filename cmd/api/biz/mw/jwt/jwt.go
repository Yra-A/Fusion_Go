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
            var req api.UserLoginRequest
            if err := c.BindAndValidate(&req); err != nil {
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
            if path == "/fusion/user/info/" {
                var req api.UserInfoRequest
                if err != c.BindAndValidate(&req) {
                    return false
                }
                if userId != req.UserID {
                    return false
                }
            } else if path == "/fusion/user/info/upload/" {
                var req api.UserInfoUploadRequest
                if err != c.BindAndValidate(&req) {
                    return false
                }
                if userId != req.UserInfo.UserID {
                    return false
                }
            } else if path == "/fusion/user/profile/:user_id" {
                var req api.UserProfileInfoRequest
                if err != c.BindAndValidate(&req) {
                    return false
                }
                if userId != req.UserID {
                    return false
                }
            } else if path == "/fusion/user/profile/upload/" {
                var req api.UserProfileUploadRequest
                if err != c.BindAndValidate(&req) {
                    return false
                }
                if userId != req.UserID || userId != req.UserProfileInfo.UserInfo.UserID {
                    return false
                }
            } else if path == "/fusion/utils/upload/img" {
                var req api.UserProfileUploadRequest
                if err != c.BindAndValidate(&req) {
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
            handler.BadResponse(c, errno.AuthorizationFailedErr)
            //path := c.FullPath()
            //if path == "/fusion/user/login/" {
            //    c.JSON(http.StatusUnauthorized, api.UserLoginResponse{
            //        StatusCode: errno.AuthorizationFailedErrCode,
            //        StatusMsg:  errno.AuthorizationFailedErr.ErrMsg,
            //    })
            //} else if path == "/fusion/user/info/" {
            //    c.JSON(http.StatusUnauthorized, api.UserInfoResponse{
            //        StatusCode: errno.AuthorizationFailedErrCode,
            //        StatusMsg:  errno.AuthorizationFailedErr.ErrMsg,
            //    })
            //} else if path == "/fusion/user/info/upload/" {
            //    c.JSON(http.StatusUnauthorized, api.UserInfoUploadResponse{
            //        StatusCode: errno.AuthorizationFailedErrCode,
            //        StatusMsg:  errno.AuthorizationFailedErr.ErrMsg,
            //    })
            //} else if path == "/fusion/user/profile/:user_id" {
            //    c.JSON(http.StatusUnauthorized, api.UserProfileInfoResponse{
            //        StatusCode: errno.AuthorizationFailedErrCode,
            //        StatusMsg:  errno.AuthorizationFailedErr.ErrMsg,
            //    })
            //} else if path == "/fusion/user/profile/upload/" {
            //    c.JSON(http.StatusUnauthorized, api.UserProfileUploadResponse{
            //        StatusCode: errno.AuthorizationFailedErrCode,
            //        StatusMsg:  errno.AuthorizationFailedErr.ErrMsg,
            //    })
            //} else if path == "/fusion/utils/upload/img" {
            //    c.JSON(http.StatusUnauthorized, api.UserProfileUploadResponse{
            //        StatusCode: errno.AuthorizationFailedErrCode,
            //        StatusMsg:  errno.AuthorizationFailedErr.ErrMsg,
            //    })
            //}
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
