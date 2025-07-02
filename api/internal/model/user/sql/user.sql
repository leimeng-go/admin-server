CREATE TABLE IF NOT EXISTS `user`  (
    `id` bigint NOT NULL AUTO_INCREMENT,
    `user_name` varchar(50) NOT NULL DEFAULT '' COMMENT '用户名',
    `password` varchar(100) NOT NULL DEFAULT '' COMMENT '密码',
    `nick_name` varchar(50) NOT NULL DEFAULT '' COMMENT '昵称',
    `avatar` varchar(255) NOT NULL DEFAULT '' COMMENT '头像',
    `role_id` bigint NOT NULL DEFAULT 0 COMMENT '角色ID',
    `status` tinyint NOT NULL DEFAULT '1' COMMENT '状态 1:正常 0:禁用 -1:删除',
    `email` varchar(100) NOT NULL DEFAULT '' COMMENT '邮箱',
    `mobile` varchar(11) NOT NULL DEFAULT '' COMMENT '手机号',
    `create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `delete_time` DATETIME DEFAULT NULL COMMENT '删除时间',
    `del_state` tinyint NOT NULL DEFAULT '0' COMMENT '删除状态 0:未删除 1:已删除',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_user_name` (`user_name`),
    UNIQUE KEY `idx_email` (`email`),
    UNIQUE KEY `idx_mobile` (`mobile`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户表'; 