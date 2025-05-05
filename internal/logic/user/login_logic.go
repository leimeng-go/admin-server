package user

import (
	"context"
	"crypto/md5"
	"encoding/hex"

	"admin-server/internal/errorx"
	"admin-server/internal/model"
	"admin-server/internal/svc"
	"admin-server/internal/types"
	"admin-server/internal/utils"

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

func (l *LoginLogic) Login(req *types.LoginReq) (map[string]interface{}, error) {
	// 查找用户
	user, err := l.svcCtx.UserModel.FindOneByUserName(l.ctx, req.Username)
	if err != nil {
		return nil, errorx.ErrUserNotFound
	}

	// 验证密码
	h := md5.New()
	h.Write([]byte(req.Password))
	passwordHash := hex.EncodeToString(h.Sum(nil))
	if passwordHash != user.Password {
		return nil, errorx.ErrInvalidParams
	}

	// 验证状态
	if user.Status != model.UserStatusNormal {
		return nil, errorx.ErrUserDisabled
	}

	token, err := utils.BuildToken(l.svcCtx.Config.Auth.AccessSecret, map[string]any{"user_id": user.Id}, l.svcCtx.Config.Auth.AccessExpire)
	if err != nil {
		return nil, errorx.ErrServerError
	}

	return map[string]interface{}{
		"token": token,
	}, nil
}
