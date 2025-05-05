CREATE TABLE `menu` (
  `id` int NOT NULL AUTO_INCREMENT,
  `menu_id` int NOT NULL COMMENT '菜单ID',
  `parent_menu_id` int DEFAULT 0 COMMENT '父菜单ID',
  `path` varchar(100) NOT NULL COMMENT '路由路径',
  `name` varchar(50) NOT NULL COMMENT '组件名',
  `component` varchar(100) DEFAULT NULL COMMENT '组件路径',
  -- meta 相关字段
  `title` varchar(50) NOT NULL COMMENT '菜单名称',
  `icon` varchar(50) DEFAULT NULL COMMENT '菜单图标',
  `show_badge` tinyint(1) DEFAULT 0 COMMENT '是否显示徽标',
  `show_text_badge` varchar(20) DEFAULT NULL COMMENT '文本徽标内容',
  `is_hide` tinyint(1) DEFAULT 0 COMMENT '是否在菜单中隐藏',
  `is_hide_tab` tinyint(1) DEFAULT 0 COMMENT '是否在标签页中隐藏',
  `link` varchar(255) DEFAULT NULL COMMENT '外部链接',
  `is_iframe` tinyint(1) DEFAULT 0 COMMENT '是否为iframe',
  `keep_alive` tinyint(1) DEFAULT 1 COMMENT '是否缓存',
  `is_in_main_container` tinyint(1) DEFAULT 0 COMMENT '是否在主容器中',
  `create_time` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `delete_time` DATETIME DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `idx_parent_id` (`parent_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='菜单表';