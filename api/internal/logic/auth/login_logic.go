package auth

import (
	"context"
	"crypto/md5"
	"encoding/hex"

	"admin-server/api/internal/errorx"
	"admin-server/api/internal/svc"
	"admin-server/api/internal/types"
	"admin-server/api/internal/pkg/utils"
	"admin-server/api/internal/model/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	userInfo,err:=l.svcCtx.UserModel.FindOneByUserName(l.ctx,req.Username)
	if err!=nil&&err==user.ErrNotFound{
		return nil,errorx.NewError(errorx.UserNotFound,"用户不存在")
	}
	if err!=nil{
		return nil,errorx.NewError(errorx.InternalServerError,"服务器内部错误")
	}
	
	// 验证用户状态
	if userInfo.Status!=user.UserStatusNormal{
		return nil,errorx.NewError(errorx.UserDisabled,"用户已禁用")
	}
	
	// 验证密码
	h:=md5.New()
	h.Write([]byte(req.Password))
	passwordHash:=hex.EncodeToString(h.Sum(nil))
	if userInfo.Password!=passwordHash{
		return nil,errorx.NewError(errorx.PasswordError,"密码错误")
	}
	
	token,err:=utils.BuildToken(l.svcCtx.Config.Auth.AccessSecret,map[string]any{"user_id":userInfo.Id},l.svcCtx.Config.Auth.AccessExpire)
	if err!=nil{
		return nil,errorx.NewError(errorx.InternalServerError,"服务器内部错误")
	}
	
	return &types.LoginResp{
		Token:token,
	},nil
}
