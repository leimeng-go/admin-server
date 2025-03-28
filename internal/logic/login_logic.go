package logic

import (
	"context"

	"github.com/leimeng-go/admin-server/internal/errorx"
	"github.com/leimeng-go/admin-server/internal/middleware"
	"github.com/leimeng-go/admin-server/internal/svc"
	"github.com/leimeng-go/admin-server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 使用用户名和密码登录系统
func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.CommonResp, err error) {
	// 参数验证
	if req.Username == "" || req.Password == "" {
		return nil, errorx.ErrInvalidParams
	}

	// 查找用户
	user, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, req.Username)
	if err != nil {
		return nil, errorx.ErrUserNotFound
	}

	// 检查用户状态
	if user.Status == 0 {
		return nil, errorx.ErrUserDisabled
	}

	// 验证密码
	if user.Password != req.Password {
		return nil, errorx.ErrPasswordError
	}

	// 生成 token
	token, err := middleware.GenerateToken(
		user.Id,
		user.Username,
		user.Role,
		l.svcCtx.Config.Auth.AccessSecret,
		l.svcCtx.Config.Auth.AccessExpire,
	)
	if err != nil {
		return nil, errorx.ErrServerError
	}

	return &types.CommonResp{
		Code:    0,
		Message: "登录成功",
		Data:    user,
		Token:   token,
	}, nil
}
