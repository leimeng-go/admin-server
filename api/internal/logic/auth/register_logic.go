package auth

import (
	"context"

	"admin-server/api/internal/model/user"
	"admin-server/api/internal/svc"
	"admin-server/api/internal/types"
	"admin-server/api/internal/errorx"
	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) error {
	if req.Password != req.ConfirmPassword {
		return errorx.ErrPasswordNotMatch
	}
	_,err:=l.svcCtx.UserModel.Insert(l.ctx,nil,&user.User{
		UserName: req.Username,
		Password: req.Password,
		// Email: req.Email,
		Status: user.UserStatusNormal,
	})
	if err != nil {
		return errorx.ErrInternalServerError
	}
	return nil
}
