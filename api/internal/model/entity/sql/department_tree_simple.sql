-- 简化版树形部门表设计（不包含触发器）
-- 支持多种树形结构查询方式：邻接表、路径枚举、嵌套集

CREATE TABLE `department_tree` (
  `id` BIGINT(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `entity_id` BIGINT(20) NOT NULL COMMENT '归属实体ID',
  `name` VARCHAR(100) NOT NULL COMMENT '部门名称',
  `code` VARCHAR(50) DEFAULT NULL COMMENT '部门编码',
  `parent_id` BIGINT(20) DEFAULT NULL COMMENT '上级部门ID，顶级部门为NULL',
  `ancestors` VARCHAR(255) DEFAULT NULL COMMENT '祖级路径，格式：1,2,3',
  `level` INT(11) NOT NULL DEFAULT 1 COMMENT '层级深度，从1开始',
  `path` VARCHAR(500) DEFAULT NULL COMMENT '完整路径，格式：/公司/技术部/开发组',
  `left_value` INT(11) DEFAULT NULL COMMENT '左值（嵌套集模型）',
  `right_value` INT(11) DEFAULT NULL COMMENT '右值（嵌套集模型）',
  `leader_id` BIGINT(20) DEFAULT NULL COMMENT '部门负责人用户ID',
  `leader_name` VARCHAR(50) DEFAULT NULL COMMENT '部门负责人姓名',
  `phone` VARCHAR(20) DEFAULT NULL COMMENT '部门联系电话',
  `email` VARCHAR(100) DEFAULT NULL COMMENT '部门邮箱',
  `description` TEXT DEFAULT NULL COMMENT '部门描述',
  `sort` INT(11) NOT NULL DEFAULT 0 COMMENT '排序',
  `status` TINYINT(1) NOT NULL DEFAULT 1 COMMENT '状态（1=正常，0=禁用）',
  `is_leaf` TINYINT(1) NOT NULL DEFAULT 1 COMMENT '是否叶子节点（1=是，0=否）',
  `create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `delete_time` DATETIME DEFAULT NULL COMMENT '删除时间',
  `del_state` TINYINT(1) NOT NULL DEFAULT 0 COMMENT '删除状态 0:未删除 1:已删除',
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
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='树形部门表，支持多种树形查询方式';

-- 创建视图：部门树形结构视图
CREATE VIEW `v_department_tree` AS
SELECT 
    d.*,
    p.name as parent_name,
    l.name as leader_name,
    (SELECT COUNT(*) FROM department_tree WHERE parent_id = d.id AND del_state = 0) as children_count
FROM department_tree d
LEFT JOIN department_tree p ON d.parent_id = p.id AND p.del_state = 0
LEFT JOIN user l ON d.leader_id = l.id AND l.del_state = 0
WHERE d.del_state = 0;

-- 插入示例数据
INSERT INTO `department_tree` (`entity_id`, `name`, `code`, `parent_id`, `level`, `path`, `left_value`, `right_value`, `description`, `sort`, `status`) VALUES
(1, '总公司', 'HQ', NULL, 1, '/总公司', 1, 14, '公司总部', 1, 1),
(1, '技术部', 'TECH', 1, 2, '/总公司/技术部', 2, 9, '负责技术研发', 1, 1),
(1, '人事部', 'HR', 1, 2, '/总公司/人事部', 10, 13, '负责人力资源', 2, 1),
(1, '开发组', 'DEV', 2, 3, '/总公司/技术部/开发组', 3, 6, '负责产品开发', 1, 1),
(1, '测试组', 'QA', 2, 3, '/总公司/技术部/测试组', 7, 8, '负责质量保证', 2, 1),
(1, '招聘组', 'RECRUIT', 3, 3, '/总公司/人事部/招聘组', 11, 12, '负责人员招聘', 1, 1);

-- 常用查询SQL示例

-- 1. 获取部门的所有子部门（使用嵌套集模型）
-- SELECT * FROM department_tree 
-- WHERE entity_id = ? AND del_state = 0 
-- AND left_value >= (SELECT left_value FROM department_tree WHERE id = ?) 
-- AND right_value <= (SELECT right_value FROM department_tree WHERE id = ?)
-- ORDER BY left_value;

-- 2. 获取部门的所有父部门（使用路径枚举）
-- SELECT * FROM department_tree 
-- WHERE entity_id = ? AND del_state = 0 
-- AND FIND_IN_SET(id, (SELECT ancestors FROM department_tree WHERE id = ?))
-- ORDER BY level;

-- 3. 获取同级部门（兄弟部门）
-- SELECT * FROM department_tree 
-- WHERE entity_id = ? AND del_state = 0 
-- AND parent_id = (SELECT parent_id FROM department_tree WHERE id = ?)
-- ORDER BY sort, id;

-- 4. 获取完整部门树
-- SELECT * FROM department_tree 
-- WHERE entity_id = ? AND del_state = 0 
-- ORDER BY left_value;

-- 5. 根据层级查询部门
-- SELECT * FROM department_tree 
-- WHERE entity_id = ? AND level = ? AND del_state = 0 
-- ORDER BY sort, id;

-- 6. 获取叶子部门
-- SELECT * FROM department_tree 
-- WHERE entity_id = ? AND is_leaf = 1 AND del_state = 0 
-- ORDER BY sort, id;

-- 7. 获取部门路径
-- SELECT path FROM department_tree WHERE id = ? AND del_state = 0;

-- 8. 统计各部门的子部门数量
-- SELECT 
--     d.id, d.name, d.level,
--     COUNT(c.id) as children_count
-- FROM department_tree d
-- LEFT JOIN department_tree c ON d.id = c.parent_id AND c.del_state = 0
-- WHERE d.entity_id = ? AND d.del_state = 0
-- GROUP BY d.id, d.name, d.level
-- ORDER BY d.left_value; 