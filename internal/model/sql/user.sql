CREATE TABLE `users` (
    `id` bigint NOT NULL AUTO_INCREMENT,
    `user_name` varchar(50) NOT NULL DEFAULT '' COMMENT '用户名',
    `password` varchar(100) NOT NULL DEFAULT '' COMMENT '密码',
    `nick_name` varchar(50) NOT NULL DEFAULT '' COMMENT '昵称',
    `avatar` varchar(255) NOT NULL DEFAULT '' COMMENT '头像',
    `role_id` bigint NOT NULL DEFAULT 0 COMMENT '角色ID',
    `status` tinyint NOT NULL DEFAULT '1' COMMENT '状态 1:正常 0:禁用 -1:删除',
    `email` varchar(100) NOT NULL DEFAULT '' COMMENT '邮箱',
    `mobile` varchar(11) NOT NULL DEFAULT '' COMMENT '手机号',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_user_name` (`user_name`),
    KEY `idx_status_deleted` (`status`, `deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户表'; 