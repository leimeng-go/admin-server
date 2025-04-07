-- 角色权限关联表
CREATE TABLE role_auth (
    `id` bigint NOT NULL AUTO_INCREMENT ,
    `role_id` BIGINT NOT NULL COMMENT '角色ID',
    `auth_id` BIGINT NOT NULL COMMENT '权限ID',
    `create_time` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `delete_time` DATETIME DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (id),
    UNIQUE KEY uk_role_auth (role_id, auth_id),
    KEY idx_auth_id (auth_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='角色权限关联表';