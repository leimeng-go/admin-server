package user

import (
	"context"

	"admin-server/internal/svc"
	"admin-server/internal/types"
	"admin-server/internal/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCurrentUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取当前登录用户的详细信息
func NewGetCurrentUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCurrentUserLogic {
	return &GetCurrentUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCurrentUserLogic) GetCurrentUser() (user *types.UserInfo, err error) {
	userID, err := utils.GetUserIdFromCtx(l.ctx)
	if err != nil {
		return nil, err
	}
	muser, err := l.svcCtx.UserModel.FindOne(l.ctx, int64(userID))
	if err != nil {
		return nil, err
	}
	return &types.UserInfo{
		ID:        muser.Id,
		Username:  muser.UserName,
		Nickname:  muser.NickName,
		Avatar:    muser.Avatar,
		Email:     muser.Email,
		Mobile:    muser.Mobile,
		RoleID:    muser.RoleId,
		Status:    muser.Status,
		CreatedAt: muser.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: muser.UpdatedAt.Format("2006-01-02 15:04:05"),
	}, nil

}
