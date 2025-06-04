CREATE TABLE `department` (
  `id` BIGINT(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `entity_id` BIGINT(20) NOT NULL COMMENT '归属实体ID',
  `name` VARCHAR(100) NOT NULL COMMENT '部门名称',
  `parent_id` BIGINT(20) DEFAULT NULL COMMENT '上级部门ID，顶级部门为NULL',
  `leader_id` BIGINT(20) DEFAULT NULL COMMENT '部门负责人用户ID',
  `sort` INT(11) NOT NULL DEFAULT 0 COMMENT '排序',
  `status` TINYINT(1) NOT NULL DEFAULT 1 COMMENT '状态（1=正常，0=禁用）',
  `create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `delete_time` DATETIME DEFAULT NULL COMMENT '删除时间',
  `del_state` tinyint NOT NULL DEFAULT '0' COMMENT '删除状态 0:未删除 1:已删除',
  PRIMARY KEY (`id`),
  KEY `idx_entity_id` (`entity_id`),
  KEY `idx_parent_id` (`parent_id`),
  KEY `idx_leader_id` (`leader_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='部门表，支持上下级归属关系，负责人为用户ID';