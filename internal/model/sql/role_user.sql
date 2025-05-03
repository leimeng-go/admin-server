CREATE TABLE role_user (
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `role_id` BIGINT NOT NULL COMMENT '角色ID',
    `user_id` BIGINT NOT NULL COMMENT '用户ID',
    `create_time` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `delete_time` DATETIME DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY uk_role_user (`role_id`, `user_id`),
    KEY idx_user_id (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='角色用户关联表';