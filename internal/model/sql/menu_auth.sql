CREATE TABLE `menu_auth` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '权限ID',
  `menu_id` bigint NOT NULL COMMENT '菜单ID',
  `title` varchar(50) NOT NULL COMMENT '权限标题',
  `auth_mark` varchar(100) NOT NULL COMMENT '权限标识',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_menu_id` (`menu_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='菜单权限表';