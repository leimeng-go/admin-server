package errorx

import (
	"fmt"
	"net/http"
)

// Code 错误码
type Code int

// 系统级错误码 (1-999)
const (
	Success       Code = 0
	ServerError   Code = 1
	InvalidParams Code = 2
	Unauthorized  Code = 3
	Forbidden     Code = 4
	NotFound      Code = 5
)

// 业务错误码 (1000-9999)
const (
	// 用户相关错误码 (1000-1999)
	UserNotFound  Code = 1000
	UserExist     Code = 1001
	PasswordError Code = 1002
	UserDisabled  Code = 1003
	TokenExpired  Code = 1004
	TokenInvalid  Code = 1005

	// 角色相关错误码 (2000-2999)
	RoleNotFound Code = 2000
	RoleExist    Code = 2001
	RoleDisabled Code = 2002

	// 权限相关错误码 (3000-3999)
	PermissionDenied Code = 3000
	PermissionExist  Code = 3001

	// 系统配置相关错误码 (4000-4999)
	ConfigNotFound Code = 4000
	ConfigExist    Code = 4001
)

// Error 自定义错误
type Error struct {
	Code    Code   // 错误码
	Message string // 错误信息
	Err     error  // 原始错误
}

func (e *Error) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("错误码: %d, 错误信息: %s, 原始错误: %s", e.Code, e.Message, e.Err.Error())
	}
	return fmt.Sprintf("错误码: %d, 错误信息: %s", e.Code, e.Message)
}

// New 创建新的错误
func New(code Code, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}

// Wrap 包装已有错误
func Wrap(err error, code Code, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

// GetHTTPCode 获取HTTP状态码
func (e *Error) GetHTTPCode() int {
	switch e.Code {
	case Success:
		return http.StatusOK
	case Unauthorized:
		return http.StatusUnauthorized
	case Forbidden:
		return http.StatusForbidden
	case NotFound:
		return http.StatusNotFound
	case InvalidParams:
		return http.StatusBadRequest
	case PasswordError:
		return http.StatusForbidden
	case UserNotFound:
		return http.StatusOK
	default:
		return http.StatusInternalServerError
	}
}

// 预定义错误
var (
	ErrSuccess       = New(Success, "成功")
	ErrServerError   = New(ServerError, "服务器内部错误")
	ErrInvalidParams = New(InvalidParams, "无效的参数")
	ErrUnauthorized  = New(Unauthorized, "未授权")
	ErrForbidden     = New(Forbidden, "禁止访问")
	ErrNotFound      = New(NotFound, "资源不存在")

	// 用户相关错误
	ErrUserNotFound  = New(UserNotFound, "用户不存在")
	ErrUserExist     = New(UserExist, "用户已存在")
	ErrPasswordError = New(PasswordError, "密码错误")
	ErrUserDisabled  = New(UserDisabled, "用户已被禁用")
	ErrTokenExpired  = New(TokenExpired, "token已过期")
	ErrTokenInvalid  = New(TokenInvalid, "无效的token")

	// 角色相关错误
	ErrRoleNotFound = New(RoleNotFound, "角色不存在")
	ErrRoleExist    = New(RoleExist, "角色已存在")
	ErrRoleDisabled = New(RoleDisabled, "角色已被禁用")

	// 权限相关错误
	ErrPermissionDenied = New(PermissionDenied, "权限不足")
	ErrPermissionExist  = New(PermissionExist, "权限已存在")

	// 系统配置相关错误
	ErrConfigNotFound = New(ConfigNotFound, "配置不存在")
	ErrConfigExist    = New(ConfigExist, "配置已存在")
)
