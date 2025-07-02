package user

import (
	"context"

	"admin-server/api/internal/model/user"
	"admin-server/api/internal/pkg/utils"
	"admin-server/api/internal/svc"
	"admin-server/api/internal/types"
	"admin-server/api/internal/errorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLogic {
	return &UserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLogic) User() (resp *types.UserInfoResp, err error) {
    userID,err:=utils.GetUserIdFromCtx(l.ctx)
    if err!=nil&&err==user.ErrNotFound{
        return nil,errorx.NewError(errorx.UserNotFound,err.Error())
    }
    user,err:=l.svcCtx.UserModel.FindOne(l.ctx,int64(userID))
    if err!=nil{
        return nil,errorx.NewError(errorx.UserNotFound,err.Error())
    }
    return &types.UserInfoResp{
        UserName: user.UserName,
        Email: user.Email,
		Mobile: user.Mobile,
		Nickname: user.NickName,
		Avatar: user.Avatar,
		// RoleID: user.RoleId,
		Status: user.Status,
		CreateTime: user.CreateTime.Format("2006-01-02 15:04:05"),
		UpdateTime: user.UpdateTime .Format("2006-01-02 15:04:05"),
    },nil
}
