package logic

import (
	"context"

	"github.com/leimeng-go/admin-server/internal/errorx"
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
		l.Logger.Infof("用户不存在: %v,username: %s", err, req.Username)
		return nil, errorx.ErrUserNotFound
	}

	// 检查用户状态
	if user.Status == 0 {
		l.Logger.Infof("用户已禁用: %v,username: %s", err, req.Username)
		return nil, errorx.ErrUserDisabled
	}

	// 验证密码
	if user.Password != req.Password {
		l.Logger.Infof("密码错误: %v,username: %s", err, req.Username)
		return nil, errorx.ErrPasswordError
	}

	// TODO: 生成 token
	token := "your-jwt-token-here"

	return &types.CommonResp{
		Code:    0,
		Message: "登录成功",
		Data:    user,
		Token:   token,
	}, nil
}
