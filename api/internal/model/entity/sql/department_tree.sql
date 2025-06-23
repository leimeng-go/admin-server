-- 树形部门表设计
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

-- 创建触发器：更新left_value和right_value（嵌套集模型）
DELIMITER $$

CREATE TRIGGER `tr_department_tree_insert` BEFORE INSERT ON `department_tree`
FOR EACH ROW
BEGIN
    DECLARE parent_left INT DEFAULT 0;
    DECLARE parent_right INT DEFAULT 0;
    DECLARE parent_level INT DEFAULT 0;
    DECLARE parent_ancestors VARCHAR(255) DEFAULT '';
    DECLARE parent_path VARCHAR(500) DEFAULT '';
    
    -- 如果有父部门，获取父部门信息
    IF NEW.parent_id IS NOT NULL THEN
        SELECT left_value, right_value, level, ancestors, path 
        INTO parent_left, parent_right, parent_level, parent_ancestors, parent_path
        FROM department_tree 
        WHERE id = NEW.parent_id AND del_state = 0;
        
        -- 设置当前部门的层级和路径
        SET NEW.level = parent_level + 1;
        SET NEW.ancestors = CASE 
            WHEN parent_ancestors = '' THEN CAST(NEW.parent_id AS CHAR)
            ELSE CONCAT(parent_ancestors, ',', NEW.parent_id)
        END;
        SET NEW.path = CASE 
            WHEN parent_path = '' THEN CONCAT('/', NEW.name)
            ELSE CONCAT(parent_path, '/', NEW.name)
        END;
        
        -- 更新所有受影响节点的left_value和right_value
        UPDATE department_tree 
        SET left_value = left_value + 2 
        WHERE left_value > parent_right AND entity_id = NEW.entity_id AND del_state = 0;
        
        UPDATE department_tree 
        SET right_value = right_value + 2 
        WHERE right_value >= parent_right AND entity_id = NEW.entity_id AND del_state = 0;
        
        -- 设置当前部门的left_value和right_value
        SET NEW.left_value = parent_right;
        SET NEW.right_value = parent_right + 1;
        
        -- 更新父部门的is_leaf状态
        UPDATE department_tree 
        SET is_leaf = 0 
        WHERE id = NEW.parent_id AND del_state = 0;
    ELSE
        -- 顶级部门
        SET NEW.level = 1;
        SET NEW.ancestors = '';
        SET NEW.path = CONCAT('/', NEW.name);
        
        -- 获取当前实体的最大right_value
        SELECT COALESCE(MAX(right_value), 0) INTO parent_right
        FROM department_tree 
        WHERE entity_id = NEW.entity_id AND del_state = 0;
        
        SET NEW.left_value = parent_right + 1;
        SET NEW.right_value = parent_right + 2;
    END IF;
END$$

CREATE TRIGGER `tr_department_tree_update` BEFORE UPDATE ON `department_tree`
FOR EACH ROW
BEGIN
    -- 如果parent_id发生变化，需要重新计算树结构
    IF OLD.parent_id != NEW.parent_id THEN
        -- 这里需要复杂的树结构调整逻辑
        -- 为了简化，建议通过应用程序来处理树结构调整
        SIGNAL SQLSTATE '45000' 
        SET MESSAGE_TEXT = 'Tree structure changes should be handled by application logic';
    END IF;
END$$

CREATE TRIGGER `tr_department_tree_delete` BEFORE UPDATE ON `department_tree`
FOR EACH ROW
BEGIN
    -- 软删除时，更新is_leaf状态
    IF NEW.del_state = 1 AND OLD.del_state = 0 THEN
        -- 检查父部门是否还有其他子部门
        IF OLD.parent_id IS NOT NULL THEN
            DECLARE sibling_count INT DEFAULT 0;
            SELECT COUNT(*) INTO sibling_count
            FROM department_tree 
            WHERE parent_id = OLD.parent_id AND id != OLD.id AND del_state = 0;
            
            IF sibling_count = 0 THEN
                UPDATE department_tree 
                SET is_leaf = 1 
                WHERE id = OLD.parent_id AND del_state = 0;
            END IF;
        END IF;
    END IF;
END$$

DELIMITER ;

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

-- 创建存储过程：获取部门的所有子部门
DELIMITER $$

CREATE PROCEDURE `sp_get_department_children`(
    IN p_entity_id BIGINT,
    IN p_department_id BIGINT,
    IN p_include_self BOOLEAN DEFAULT FALSE
)
BEGIN
    DECLARE v_left_value INT;
    DECLARE v_right_value INT;
    
    -- 获取目标部门的左右值
    SELECT left_value, right_value 
    INTO v_left_value, v_right_value
    FROM department_tree 
    WHERE id = p_department_id AND entity_id = p_entity_id AND del_state = 0;
    
    -- 查询所有子部门
    SELECT * FROM department_tree 
    WHERE entity_id = p_entity_id 
    AND del_state = 0
    AND left_value >= v_left_value 
    AND right_value <= v_right_value
    AND (p_include_self = TRUE OR id != p_department_id)
    ORDER BY left_value;
END$$

-- 创建存储过程：获取部门的所有父部门
CREATE PROCEDURE `sp_get_department_parents`(
    IN p_entity_id BIGINT,
    IN p_department_id BIGINT,
    IN p_include_self BOOLEAN DEFAULT FALSE
)
BEGIN
    DECLARE v_ancestors VARCHAR(255);
    
    -- 获取目标部门的祖先路径
    SELECT ancestors 
    INTO v_ancestors
    FROM department_tree 
    WHERE id = p_department_id AND entity_id = p_entity_id AND del_state = 0;
    
    -- 查询所有父部门
    IF v_ancestors IS NOT NULL AND v_ancestors != '' THEN
        SELECT * FROM department_tree 
        WHERE entity_id = p_entity_id 
        AND del_state = 0
        AND FIND_IN_SET(id, v_ancestors)
        AND (p_include_self = TRUE OR id != p_department_id)
        ORDER BY level;
    END IF;
END$$

DELIMITER ;

-- 插入示例数据
INSERT INTO `department_tree` (`entity_id`, `name`, `code`, `parent_id`, `level`, `path`, `left_value`, `right_value`, `description`, `sort`, `status`) VALUES
(1, '总公司', 'HQ', NULL, 1, '/总公司', 1, 14, '公司总部', 1, 1),
(1, '技术部', 'TECH', 1, 2, '/总公司/技术部', 2, 9, '负责技术研发', 1, 1),
(1, '人事部', 'HR', 1, 2, '/总公司/人事部', 10, 13, '负责人力资源', 2, 1),
(1, '开发组', 'DEV', 2, 3, '/总公司/技术部/开发组', 3, 6, '负责产品开发', 1, 1),
(1, '测试组', 'QA', 2, 3, '/总公司/技术部/测试组', 7, 8, '负责质量保证', 2, 1),
(1, '招聘组', 'RECRUIT', 3, 3, '/总公司/人事部/招聘组', 11, 12, '负责人员招聘', 1, 1); 