package errno

import (
	"errors"
	"fmt"
)

const (
	SuccessCode                    = 0
	FailCode                       = 10000
	ServiceErrCode                 = 10001
	ParamErrCode                   = 10002
	AuthorizationFailedErrCode     = 10003
	InvalidCredentialsErrCode      = 10004
	UserNotFoundErrCode            = 10005
	UserinfoNotSetCode             = 10006
	UserAlreadyExistErrCode        = 10007
	EmptyUsernameOrPasswordErrCode = 10008
	ContestNotExistErrCode         = 10009
	ArticleNotExistErrCode         = 10010
)

type ErrNo struct {
	ErrCode int32
	ErrMsg  string
}

func (e ErrNo) Error() string {
	return fmt.Sprintf("err_code=%d, err_msg=%s", e.ErrCode, e.ErrMsg)
}

func NewErrNo(code int32, msg string) ErrNo {
	return ErrNo{code, msg}
}

func (e ErrNo) WithMessage(msg string) ErrNo {
	e.ErrMsg = msg
	return e
}

var (
	Success                    = NewErrNo(SuccessCode, "成功")
	ServiceErr                 = NewErrNo(ServiceErrCode, "服务未能成功启动")
	ParamErr                   = NewErrNo(ParamErrCode, "参数错误")
	Fail                       = NewErrNo(FailCode, "出现失败")
	AuthorizationFailedErr     = NewErrNo(AuthorizationFailedErrCode, "授权失败")
	InvalidCredentialsErr      = NewErrNo(InvalidCredentialsErrCode, "用户名或密码错误")
	UserNotExistErr            = NewErrNo(UserNotFoundErrCode, "用户不存在")
	UserinfoNotSetErr          = NewErrNo(UserinfoNotSetCode, "用户信息未设置")
	UserAlreadyExistErr        = NewErrNo(UserAlreadyExistErrCode, "用户已经存在")
	EmptyUsernameOrPasswordErr = NewErrNo(EmptyUsernameOrPasswordErrCode, "用户名或密码为空")
	ContestNotExistErr         = NewErrNo(ContestNotExistErrCode, "赛事不存在")
	ArticleNotExistErr         = NewErrNo(ArticleNotExistErrCode, "文章不存在")
)

// ConvertErr convert error to Errno
func ConvertErr(err error) ErrNo {
	Err := ErrNo{}
	if errors.As(err, &Err) {
		return Err
	}

	s := ServiceErr
	s.ErrMsg = err.Error()
	return s
}
