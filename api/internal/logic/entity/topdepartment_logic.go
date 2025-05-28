package entity

import (
	"context"

	"admin-server/api/internal/pkg/utils"
	"admin-server/api/internal/svc"
	"admin-server/api/internal/types"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/logx"
)

type TopdepartmentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTopdepartmentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TopdepartmentLogic {
	return &TopdepartmentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TopdepartmentLogic) Topdepartment(req *types.TopDepartmentReq) (resp *types.TopDepartmentResp, err error) {
	userID,err:=utils.GetUserIdFromCtx(l.ctx)
	if err!=nil{
		return nil,err
	}
	entityUser,err:=l.svcCtx.EntityUserModel.FindOne(l.ctx,userID)
	if err!=nil{
		return nil,err
	}
	
	departmentQuery:=l.svcCtx.DepartmentModel.SelectBuilder().
	Where(squirrel.Eq{
		"entity_id":entityUser.EntityId,
	})
	if req.ParentID==0{
		departmentQuery=departmentQuery.Where(squirrel.Eq{
			"parent_id":nil,
		})
	}else{
		departmentQuery=departmentQuery.Where(squirrel.Eq{
			"parent_id":req.ParentID,
		})
	}
	
	departmentList,err:=l.svcCtx.DepartmentModel.FindAll(l.ctx,departmentQuery,"id desc")
	if err!=nil{
		return nil,err
	}
	resp=new(types.TopDepartmentResp)
	resp.List=make([]types.TopDepartment,0,len(departmentList))
	for _,v:=range departmentList{
		resp.List=append(resp.List,types.TopDepartment{
			ID:v.Id,
			Name:v.Name,
		})
	}
	return
}
