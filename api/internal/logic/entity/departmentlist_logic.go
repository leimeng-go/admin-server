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
	userID, err := utils.GetUserIdFromCtx(l.ctx)
	if err != nil {
		return nil, err
	}

	entityUserQuery := l.svcCtx.EntityUserModel.SelectBuilder().Where(
		squirrel.Eq{"user_id": userID},
	)

	entityUsers, err := l.svcCtx.EntityUserModel.FindAll(l.ctx, entityUserQuery, "id")
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	entityIds := make([]int64, 0, len(entityUsers))
	for _, v := range entityUsers {
		entityIds = append(entityIds, v.EntityId)
	}
	// 查询一级部门（parent_id 为 0 或 NULL）
	departmentQuery := l.svcCtx.DepartmentModel.SelectBuilder().Where(
		squirrel.Eq{"entity_id": entityIds})
	// 兼容 parent_id 为 0 或 NULL
	departmentQuery = departmentQuery.Where(squirrel.Or{
		squirrel.Eq{"parent_id": 0},
		squirrel.Eq{"parent_id": nil},
	})
	if req.Keyword != "" {
		departmentQuery = departmentQuery.Where(squirrel.Like{
			"name": "%" + req.Keyword + "%",
		})
	}
	departmentList, total, err := l.svcCtx.DepartmentModel.FindPageListByPageWithTotal(l.ctx, departmentQuery, req.Page, req.PageSize, "sort")
	if err != nil {
		return nil, err
	}
	resp = &types.DepartmentListResp{
		Total: total,
		List:  make([]types.Department, 0, len(departmentList)),
	}
	for _, v := range departmentList {
		children, err := l.getDepartmentChildren(v.Id)
		if err != nil {
			return nil, err
		}
		resp.List = append(resp.List, types.Department{
			ID:         v.Id,
			Name:       v.Name,
			Sort:       v.Sort,
			Status:     v.Status,
			CreateTime: carbon.CreateFromStdTime(v.CreateTime).ToDateTimeString(),
			UpdateTime: carbon.CreateFromStdTime(v.UpdateTime).ToDateTimeString(),
			Children:   children,
		})
	}
	return
}

// 获取某部门的所有子部门（递归）
func (l *DepartmentlistLogic) getDepartmentChildren(departmentId int64) ([]types.Department, error) {
	query := l.svcCtx.DepartmentModel.SelectBuilder().Where(
		squirrel.Eq{"parent_id": departmentId},
	)
	children, err := l.svcCtx.DepartmentModel.FindAll(l.ctx, query, "sort")
	if err != nil {
		return nil, err
	}
	result := make([]types.Department, 0, len(children))
	for _, v := range children {
		// 递归查找下级部门
		grandChildren, err := l.getDepartmentChildren(v.Id)
		if err != nil {
			return nil, err
		}
		result = append(result, types.Department{
			ID:         v.Id,
			Name:       v.Name,
			Sort:       v.Sort,
			Status:     v.Status,
			CreateTime: carbon.CreateFromStdTime(v.CreateTime).ToDateTimeString(),
			UpdateTime: carbon.CreateFromStdTime(v.UpdateTime).ToDateTimeString(),
			Children:   grandChildren,
		})
	}
	return result, nil
}
