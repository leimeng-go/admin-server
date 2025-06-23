package entity

import (
	"context"
	"database/sql"
	"fmt"
)

// DepartmentTreeNode 树形节点结构
type DepartmentTreeNode struct {
	*DepartmentTree
	Children []*DepartmentTreeNode
}

// BuildDepartmentTree 构建树形结构
func BuildDepartmentTree(departments []*DepartmentTree) []*DepartmentTreeNode {
	deptMap := make(map[int64]*DepartmentTreeNode)
	var roots []*DepartmentTreeNode

	// 创建节点映射
	for _, dept := range departments {
		node := &DepartmentTreeNode{
			DepartmentTree: dept,
			Children:       make([]*DepartmentTreeNode, 0),
		}
		deptMap[dept.Id] = node
	}

	// 构建树形结构
	for _, dept := range departments {
		node := deptMap[dept.Id]
		if dept.ParentId.Valid {
			if parent, exists := deptMap[dept.ParentId.Int64]; exists {
				parent.Children = append(parent.Children, node)
			}
		} else {
			roots = append(roots, node)
		}
	}

	return roots
}

// PrintDepartmentTree 打印树形结构（用于调试）
func PrintDepartmentTree(nodes []*DepartmentTreeNode, level int) {
	for _, node := range nodes {
		indent := ""
		for i := 0; i < level; i++ {
			indent += "  "
		}
		fmt.Printf("%s├─ %s (ID: %d, Level: %d)\n", indent, node.Name, node.Id, node.Level)
		if len(node.Children) > 0 {
			PrintDepartmentTree(node.Children, level+1)
		}
	}
}

// ExampleUsage 使用示例
func ExampleUsage(ctx context.Context, model DepartmentTreeModel) {
	entityId := int64(1)

	// 1. 创建顶级部门
	rootDept := &DepartmentTree{
		EntityId:    entityId,
		Name:        "总公司",
		Code:        sql.NullString{String: "HQ", Valid: true},
		ParentId:    sql.NullInt64{Valid: false}, // 顶级部门
		Description: sql.NullString{String: "公司总部", Valid: true},
		Sort:        1,
		Status:      1,
	}
	_, err := model.Insert(ctx, nil, rootDept)
	if err != nil {
		fmt.Printf("创建顶级部门失败: %v\n", err)
		return
	}

	// 2. 创建子部门
	techDept := &DepartmentTree{
		EntityId:    entityId,
		Name:        "技术部",
		Code:        sql.NullString{String: "TECH", Valid: true},
		ParentId:    sql.NullInt64{Int64: rootDept.Id, Valid: true},
		Description: sql.NullString{String: "负责技术研发", Valid: true},
		Sort:        1,
		Status:      1,
	}
	_, err = model.Insert(ctx, nil, techDept)
	if err != nil {
		fmt.Printf("创建技术部失败: %v\n", err)
		return
	}

	// 3. 创建孙部门
	devDept := &DepartmentTree{
		EntityId:    entityId,
		Name:        "开发组",
		Code:        sql.NullString{String: "DEV", Valid: true},
		ParentId:    sql.NullInt64{Int64: techDept.Id, Valid: true},
		Description: sql.NullString{String: "负责产品开发", Valid: true},
		Sort:        1,
		Status:      1,
	}
	_, err = model.Insert(ctx, nil, devDept)
	if err != nil {
		fmt.Printf("创建开发组失败: %v\n", err)
		return
	}

	// 4. 查询完整部门树
	tree, err := model.FindTree(ctx, entityId)
	if err != nil {
		fmt.Printf("查询部门树失败: %v\n", err)
		return
	}

	// 5. 构建树形结构
	treeNodes := BuildDepartmentTree(tree)
	fmt.Println("部门树结构:")
	PrintDepartmentTree(treeNodes, 0)

	// 6. 查询技术部的所有子部门
	children, err := model.FindAllChildren(ctx, entityId, techDept.Id, false)
	if err != nil {
		fmt.Printf("查询子部门失败: %v\n", err)
		return
	}
	fmt.Printf("\n技术部的子部门数量: %d\n", len(children))

	// 7. 查询开发组的所有父部门
	parents, err := model.FindParents(ctx, entityId, devDept.Id, false)
	if err != nil {
		fmt.Printf("查询父部门失败: %v\n", err)
		return
	}
	fmt.Printf("开发组的父部门数量: %d\n", len(parents))

	// 8. 查询同级部门（兄弟部门）
	siblings, err := model.FindSiblings(ctx, entityId, devDept.Id, false)
	if err != nil {
		fmt.Printf("查询兄弟部门失败: %v\n", err)
		return
	}
	fmt.Printf("开发组的兄弟部门数量: %d\n", len(siblings))

	// 9. 根据层级查询
	level2Depts, err := model.FindByLevel(ctx, entityId, 2)
	if err != nil {
		fmt.Printf("查询第2层级部门失败: %v\n", err)
		return
	}
	fmt.Printf("第2层级部门数量: %d\n", len(level2Depts))
}

// ValidateDepartmentMove 验证部门移动的合法性
func ValidateDepartmentMove(ctx context.Context, model DepartmentTreeModel, entityId, departmentId, newParentId int64) error {
	// 1. 验证目标父部门是否存在
	if newParentId > 0 {
		_, err := model.FindOne(ctx, newParentId)
		if err != nil {
			return fmt.Errorf("目标父部门不存在: %v", err)
		}
	}

	// 2. 验证不能移动到自己的子部门
	children, err := model.FindAllChildren(ctx, entityId, departmentId, false)
	if err != nil {
		return fmt.Errorf("查询子部门失败: %v", err)
	}
	for _, child := range children {
		if child.Id == newParentId {
			return fmt.Errorf("不能将部门移动到自己的子部门")
		}
	}

	// 3. 验证不能移动到自己的父部门
	if newParentId == departmentId {
		return fmt.Errorf("不能将部门移动到自己的父部门")
	}

	return nil
}

// GetDepartmentPath 获取部门的完整路径
func GetDepartmentPath(ctx context.Context, model DepartmentTreeModel, entityId, departmentId int64) (string, error) {
	dept, err := model.FindOne(ctx, departmentId)
	if err != nil {
		return "", err
	}

	if dept.Path.Valid {
		return dept.Path.String, nil
	}

	// 如果path字段为空，通过ancestors构建路径
	if dept.Ancestors.Valid && dept.Ancestors.String != "" {
		// 这里可以通过ancestors字段构建路径
		// 简化处理，直接返回name
		return "/" + dept.Name, nil
	}

	return "/" + dept.Name, nil
}

// GetDepartmentLevel 获取部门的层级深度
func GetDepartmentLevel(ctx context.Context, model DepartmentTreeModel, entityId, departmentId int64) (int, error) {
	dept, err := model.FindOne(ctx, departmentId)
	if err != nil {
		return 0, err
	}
	return int(dept.Level), nil
}

// IsLeafDepartment 判断是否为叶子部门
func IsLeafDepartment(ctx context.Context, model DepartmentTreeModel, entityId, departmentId int64) (bool, error) {
	dept, err := model.FindOne(ctx, departmentId)
	if err != nil {
		return false, err
	}
	return dept.IsLeaf == 1, nil
}
