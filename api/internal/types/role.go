// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.4

package types

type AddRoleReq struct {
	Name        string   `json:"name"`        // 角色名称
	Policies    []Policy `json:"policies"`    // 策略
	Description string   `json:"description"` // 角色描述
	Status      int64    `json:"status"`      // 状态
}

type DeleteRoleReq struct {
	ID int64 `form:"id"` // 角色ID
}

type RoleListReq struct {
	Page     int64 `form:"page"`     // 页码
	PageSize int64 `form:"pageSize"` // 每页条数
}

type RoleListResp struct {
	List  []Role `json:"list"`
	Total int64  `json:"total"`
}

type UpdateRoleReq struct {
	ID          int64  `json:"id"`          // 角色ID
	Name        string `json:"name"`        // 角色名称
	Description string `json:"description"` // 角色描述
	Status      int64  `json:"status"`      // 状态
}
