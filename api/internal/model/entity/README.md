# 树形部门表设计总结

## 设计概述

我为你的管理系统设计了一个完善的树形部门表结构，支持多种树形查询方式，满足不同场景的性能需求。

## 文件说明

### 1. SQL文件
- `department_tree.sql` - 完整版，包含触发器和存储过程
- `department_tree_simple.sql` - 简化版，不包含触发器，适合所有环境

### 2. Go模型文件
- `department_tree_model.go` - 完整的Go模型实现
- `department_tree_example.go` - 使用示例和工具函数
- `department_tree_usage.md` - 详细使用说明

## 表结构特点

### 核心字段设计
```sql
CREATE TABLE `department_tree` (
  `id` BIGINT(20) NOT NULL AUTO_INCREMENT,
  `entity_id` BIGINT(20) NOT NULL,           -- 多租户支持
  `name` VARCHAR(100) NOT NULL,              -- 部门名称
  `code` VARCHAR(50) DEFAULT NULL,           -- 部门编码（实体内唯一）
  `parent_id` BIGINT(20) DEFAULT NULL,       -- 上级部门ID
  `ancestors` VARCHAR(255) DEFAULT NULL,     -- 祖级路径
  `level` INT(11) NOT NULL DEFAULT 1,        -- 层级深度
  `path` VARCHAR(500) DEFAULT NULL,          -- 完整路径
  `left_value` INT(11) DEFAULT NULL,         -- 嵌套集左值
  `right_value` INT(11) DEFAULT NULL,        -- 嵌套集右值
  `is_leaf` TINYINT(1) NOT NULL DEFAULT 1,   -- 是否叶子节点
  -- 其他业务字段...
)
```

### 三种树形模型结合
1. **邻接表模型** (`parent_id`) - 简单直观，适合基本操作
2. **路径枚举** (`ancestors`, `path`) - 查询父级路径高效
3. **嵌套集模型** (`left_value`, `right_value`) - 查询子级结构高效

## 主要功能

### 1. 基础CRUD操作
- 创建部门（自动计算树形字段）
- 查询部门信息
- 更新部门信息
- 软删除部门

### 2. 树形查询操作
- 获取直接子部门
- 获取所有子部门（包括子部门的子部门）
- 获取所有父部门
- 获取兄弟部门
- 根据层级查询
- 获取完整部门树

### 3. 业务功能
- 部门移动（改变父部门）
- 部门路径获取
- 层级深度计算
- 叶子节点判断

## 性能优化

### 索引设计
```sql
PRIMARY KEY (`id`),
UNIQUE KEY `uk_entity_code` (`entity_id`, `code`),
KEY `idx_entity_id` (`entity_id`),
KEY `idx_parent_id` (`parent_id`),
KEY `idx_ancestors` (`ancestors`),
KEY `idx_level` (`level`),
KEY `idx_left_right` (`left_value`, `right_value`),
KEY `idx_leader_id` (`leader_id`),
KEY `idx_status` (`status`),
KEY `idx_sort` (`sort`)
```

### 查询性能
- **子部门查询**: 使用嵌套集模型，O(1)复杂度
- **父部门查询**: 使用路径枚举，O(n)复杂度
- **同级查询**: 使用邻接表模型，简单高效
- **完整树查询**: 按left_value排序，支持递归构建

## 使用示例

### 创建部门
```go
dept := &DepartmentTree{
    EntityId:    1,
    Name:        "技术部",
    Code:        sql.NullString{String: "TECH", Valid: true},
    ParentId:    sql.NullInt64{Int64: 1, Valid: true},
    Description: sql.NullString{String: "负责技术研发", Valid: true},
    Sort:        1,
    Status:      1,
}
result, err := model.Insert(ctx, nil, dept)
```

### 查询部门树
```go
// 获取完整部门树
tree, err := model.FindTree(ctx, entityId)

// 构建树形结构
treeNodes := BuildDepartmentTree(tree)

// 打印树形结构
PrintDepartmentTree(treeNodes, 0)
```

### 常用查询
```go
// 获取子部门
children, err := model.FindChildren(ctx, entityId, parentId)

// 获取所有子部门
allChildren, err := model.FindAllChildren(ctx, entityId, departmentId, false)

// 获取父部门
parents, err := model.FindParents(ctx, entityId, departmentId, false)

// 获取兄弟部门
siblings, err := model.FindSiblings(ctx, entityId, departmentId, false)
```

## 扩展功能

### 1. 部门权限控制
基于树形结构可以实现：
- 用户只能访问自己所在部门及其子部门
- 管理员可以管理指定部门及其子部门

### 2. 组织架构图
可以基于树形数据生成：
- 层级展示
- 折叠展开
- 拖拽调整

### 3. 部门统计
基于树形结构可以轻松实现：
- 各部门人员统计
- 部门层级统计
- 部门预算分配

## 注意事项

1. **并发安全**: 插入和删除操作会修改树形结构，需要适当的锁机制
2. **数据一致性**: 复杂的树结构调整建议在事务中执行
3. **缓存策略**: 树形结构变化时，需要清理相关缓存
4. **软删除**: 支持软删除，删除时会自动更新父部门的is_leaf状态

## 部署建议

1. **开发环境**: 使用简化版SQL文件，不包含触发器
2. **生产环境**: 根据数据库支持情况选择完整版或简化版
3. **性能调优**: 根据实际查询模式调整索引策略
4. **监控告警**: 监控树形结构的完整性和一致性

这个设计既保证了功能的完整性，又考虑了性能优化，可以满足大多数企业级应用的部门管理需求。 