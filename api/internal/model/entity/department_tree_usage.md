# 树形部门表设计说明

## 概述

这个树形部门表设计支持多种树形结构查询方式，包括：
- **邻接表模型**：通过 `parent_id` 字段建立父子关系
- **路径枚举**：通过 `ancestors` 和 `path` 字段记录完整路径
- **嵌套集模型**：通过 `left_value` 和 `right_value` 字段支持高效查询

## 表结构特点

### 核心字段
- `id`: 主键
- `entity_id`: 归属实体ID（支持多租户）
- `name`: 部门名称
- `code`: 部门编码（实体内唯一）
- `parent_id`: 上级部门ID
- `ancestors`: 祖级路径（格式：1,2,3）
- `level`: 层级深度
- `path`: 完整路径（格式：/公司/技术部/开发组）
- `left_value`/`right_value`: 嵌套集模型的左右值
- `is_leaf`: 是否叶子节点

### 业务字段
- `leader_id`/`leader_name`: 部门负责人
- `phone`/`email`: 部门联系方式
- `description`: 部门描述
- `sort`: 排序
- `status`: 状态

## 查询方式

### 1. 获取直接子部门
```go
children, err := model.FindChildren(ctx, entityId, parentId)
```

### 2. 获取所有子部门（包括子部门的子部门）
```go
allChildren, err := model.FindAllChildren(ctx, entityId, departmentId, false)
```

### 3. 获取所有父部门
```go
parents, err := model.FindParents(ctx, entityId, departmentId, false)
```

### 4. 获取完整部门树
```go
tree, err := model.FindTree(ctx, entityId)
```

### 5. 根据层级查询
```go
level2Depts, err := model.FindByLevel(ctx, entityId, 2)
```

### 6. 获取兄弟部门
```go
siblings, err := model.FindSiblings(ctx, entityId, departmentId, false)
```

## 自动维护机制

### 触发器
- **插入触发器**: 自动计算 `level`、`ancestors`、`path`、`left_value`、`right_value`
- **更新触发器**: 防止直接修改 `parent_id`（需要通过应用程序逻辑处理）
- **删除触发器**: 自动更新父部门的 `is_leaf` 状态

### 存储过程
- `sp_get_department_children`: 获取所有子部门
- `sp_get_department_parents`: 获取所有父部门

## 使用示例

### 创建部门
```go
dept := &DepartmentTree{
    EntityId:    1,
    Name:        "技术部",
    Code:        sql.NullString{String: "TECH", Valid: true},
    ParentId:    sql.NullInt64{Int64: 1, Valid: true}, // 父部门ID
    Description: sql.NullString{String: "负责技术研发", Valid: true},
    Sort:        1,
    Status:      1,
}

result, err := model.Insert(ctx, nil, dept)
```

### 构建树形结构
```go
func BuildDepartmentTree(departments []*DepartmentTree) []*DepartmentTreeNode {
    deptMap := make(map[int64]*DepartmentTreeNode)
    var roots []*DepartmentTreeNode
    
    // 创建节点映射
    for _, dept := range departments {
        node := &DepartmentTreeNode{
            Department: dept,
            Children:   make([]*DepartmentTreeNode, 0),
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

type DepartmentTreeNode struct {
    *DepartmentTree
    Children []*DepartmentTreeNode
}
```

### 移动部门
```go
func MoveDepartment(ctx context.Context, model DepartmentTreeModel, departmentId, newParentId int64) error {
    return model.Trans(ctx, func(ctx context.Context, session sqlx.Session) error {
        // 1. 验证目标父部门是否存在
        if newParentId > 0 {
            _, err := model.FindOne(ctx, newParentId)
            if err != nil {
                return err
            }
        }
        
        // 2. 验证不能移动到自己的子部门
        children, err := model.FindAllChildren(ctx, entityId, departmentId, false)
        if err != nil {
            return err
        }
        for _, child := range children {
            if child.Id == newParentId {
                return errors.New("cannot move department to its own child")
            }
        }
        
        // 3. 执行移动操作（需要重新计算整个子树的结构）
        // 这里需要复杂的逻辑来重新计算 left_value 和 right_value
        // 建议使用专门的存储过程或应用程序逻辑
        
        return nil
    })
}
```

## 性能优化

### 索引设计
- `idx_entity_id`: 按实体查询
- `idx_parent_id`: 按父部门查询
- `idx_ancestors`: 按祖先路径查询
- `idx_left_right`: 嵌套集查询
- `idx_level`: 按层级查询

### 查询优化
1. **获取子部门**: 使用嵌套集模型，性能最佳
2. **获取父部门**: 使用路径枚举，性能良好
3. **获取兄弟部门**: 使用邻接表模型，简单高效
4. **完整树查询**: 按 `left_value` 排序，支持递归构建

## 注意事项

1. **并发安全**: 插入和删除操作会修改 `left_value` 和 `right_value`，需要适当的锁机制
2. **数据一致性**: 复杂的树结构调整建议在事务中执行
3. **缓存策略**: 树形结构变化时，需要清理相关缓存
4. **软删除**: 支持软删除，删除时会自动更新父部门的 `is_leaf` 状态

## 扩展功能

### 部门权限
可以基于树形结构实现部门权限控制：
- 用户只能访问自己所在部门及其子部门
- 管理员可以管理指定部门及其子部门

### 部门统计
基于树形结构可以轻松实现：
- 各部门人员统计
- 部门层级统计
- 部门预算分配

### 组织架构图
可以基于树形数据生成组织架构图，支持：
- 层级展示
- 折叠展开
- 拖拽调整 