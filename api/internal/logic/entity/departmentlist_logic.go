package entity

import (
	"context"
	"database/sql"

	"admin-server/api/internal/pkg/utils"
	"admin-server/api/internal/svc"
	"admin-server/api/internal/types"

	"github.com/Masterminds/squirrel"
	"github.com/golang-module/carbon/v2"
	"github.com/zeromicro/go-zero/core/logx"
)

type DepartmentlistLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDepartmentlistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DepartmentlistLogic {
	return &DepartmentlistLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DepartmentlistLogic) Departmentlist(req *types.DepartmentListReq) (resp *types.DepartmentListResp, err error) {
	userID,err:=utils.GetUserIdFromCtx(l.ctx)
	if err!=nil{
		return nil,err
	}
	
	entityUserQuery:=l.svcCtx.EntityUserModel.SelectBuilder().Where(
		squirrel.Eq{"user_id":userID},
	)
	entityUsers,err:=l.svcCtx.EntityUserModel.FindAll(l.ctx,entityUserQuery,"id")
	if err!=nil&&err!=sql.ErrNoRows{
		return nil,err
	}
    entityIds:=make([]int64,0,len(entityUsers))
	for _,v:=range entityUsers{
		entityIds=append(entityIds,v.EntityId)
	}
	
	departmentQuery:=l.svcCtx.DepartmentModel.SelectBuilder().Where(
		squirrel.Eq{"entity_id":entityIds},)	
	offset,limit:=utils.Page(req.Page,req.PageSize)
	departmentList,total,err:=l.svcCtx.DepartmentModel.FindPageListByPageWithTotal(l.ctx,departmentQuery,offset,limit,"sort")
	if err!=nil{
		return nil,err
	}
	resp=&types.DepartmentListResp{
		Total:total,
		List:make([]types.Department,0,len(departmentList)),
	}
     
	for _,v:=range departmentList{
		resp.List=append(resp.List,types.Department{
			ID: v.Id,
			Name: v.Name,
			Sort: v.Sort,
			Status: v.Status,
			CreatedTime: carbon.CreateFromStdTime(v.CreateTime).ToDateTimeString(),
			UpdatedTime: carbon.CreateFromStdTime(v.UpdateTime).ToDateTimeString(),
			Children: make([]types.Department,0),
		})
	}
	return
}
