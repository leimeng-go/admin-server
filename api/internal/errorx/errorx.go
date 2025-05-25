package errorx 

import(
	"fmt"
	"net/http"
)

// Code 错误码
type Code int 

const(
	Success Code = 0
)

// 业务错误码
const (
   // 用户相关错误码 (1000-1999)
	UserNotFound  Code = 1000
	UserExist     Code = 1001
	PasswordError Code = 1002
	UserDisabled  Code = 1003
	TokenExpired  Code = 1004
	TokenInvalid  Code = 1005


	// 实体相关错误码 (1100-1199)
	EntityNotFound Code = 1100

	// 系统相关错误码 (2000-2999)
	InternalServerError Code = 9999
	
)

type Error struct {
	Code    Code
	Message string
	Err     error
}

func (e *Error) Error() string {
	if e.Err!=nil{
		return fmt.Sprintf("错误码: %d, 错误信息: %s, 错误原因: %s", e.Code, e.Message, e.Err)
	}
	return fmt.Sprintf("错误码: %d, 错误信息: %s", e.Code, e.Message)
}

func (e *Error) Unwrap() error {
	return e.Err
}

func NewError(code Code, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}

func Wrap(err error, code Code, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

func (e *Error) GetHTTPCode() int {
	switch e.Code {
	case Success:
		return http.StatusOK
	case UserDisabled:
		return http.StatusForbidden
	case TokenExpired, TokenInvalid:
		return http.StatusUnauthorized
	case InternalServerError:
		return http.StatusInternalServerError
	
	}
	return http.StatusOK
}

// 预定义错误
var(
	ErrSuccess = NewError(Success, "成功")
	ErrUserNotFound = NewError(UserNotFound, "用户不存在")
	ErrUserExist = NewError(UserExist, "用户已存在")
	ErrPasswordError = NewError(PasswordError, "密码错误")
	ErrUserDisabled = NewError(UserDisabled, "用户禁用")
	ErrTokenExpired = NewError(TokenExpired, "token过期")
	ErrTokenInvalid = NewError(TokenInvalid, "token无效")

	ErrEntityNotFound = NewError(EntityNotFound, "实体不存在")
	ErrInternalServerError = NewError(InternalServerError, "服务器内部错误")
)
