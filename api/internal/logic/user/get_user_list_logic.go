package user

import (
	"context"

	"admin-server/api/internal/svc"
	"admin-server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserListLogic {
	return &GetUserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserListLogic) GetUserList(req *types.GetUserListReq) (resp *types.GetUserListResp, err error) {
	userList,total,err:=l.svcCtx.UserModel.FindPageListByPageWithTotal(l.ctx,l.svcCtx.UserModel.SelectBuilder(),req.Current,req.Size,"")
	if err!=nil{
		return nil,err
	}
	resp=new(types.GetUserListResp)
	for _,user:=range userList{
		resp.List=append(resp.List,types.UserInfoResp{
			UserName: user.UserName,
			Email: user.Email,
			Status: user.Status,
			CreateTime: user.CreateTime.Format("2006-01-02 15:04:05"),
			UpdateTime: user.UpdateTime.Format("2006-01-02 15:04:05"),
		})
	}
	resp.Current=req.Current
	resp.Size=req.Size
	resp.Total=total
	return
}
