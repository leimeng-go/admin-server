-- 创建用户表
CREATE TABLE IF NOT EXISTS users (
    id BIGINT NOT NULL AUTO_INCREMENT,
    username VARCHAR(50) NOT NULL,
    password VARCHAR(100) NOT NULL,
    nickname VARCHAR(50) NOT NULL,
    avatar VARCHAR(255) NOT NULL,
    email VARCHAR(100) NOT NULL,
    mobile VARCHAR(20) NOT NULL,
    role VARCHAR(20) NOT NULL DEFAULT 'user',
    status TINYINT NOT NULL DEFAULT 1,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    PRIMARY KEY (id),
    UNIQUE KEY uk_username (username),
    UNIQUE KEY uk_email (email),
    UNIQUE KEY uk_mobile (mobile),
    KEY idx_status (status)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 插入管理员用户
INSERT INTO users (username, password, nickname, avatar, email, mobile, role, status) VALUES
('admin', 'e10adc3949ba59abbe56e057f20f883e', '系统管理员', 'https://avatars.githubusercontent.com/u/1?v=4', 'admin@example.com', '13800000001', 'admin', 1),
('superadmin', 'e10adc3949ba59abbe56e057f20f883e', '超级管理员', 'https://avatars.githubusercontent.com/u/2?v=4', 'superadmin@example.com', '13800000002', 'admin', 1);

-- 插入普通用户
INSERT INTO users (username, password, nickname, avatar, email, mobile, role, status) VALUES
('user1', 'e10adc3949ba59abbe56e057f20f883e', '张三', 'https://avatars.githubusercontent.com/u/3?v=4', 'zhangsan@example.com', '13800000003', 'user', 1),
('user2', 'e10adc3949ba59abbe56e057f20f883e', '李四', 'https://avatars.githubusercontent.com/u/4?v=4', 'lisi@example.com', '13800000004', 'user', 1),
('user3', 'e10adc3949ba59abbe56e057f20f883e', '王五', 'https://avatars.githubusercontent.com/u/5?v=4', 'wangwu@example.com', '13800000005', 'user', 1),
('user4', 'e10adc3949ba59abbe56e057f20f883e', '赵六', 'https://avatars.githubusercontent.com/u/6?v=4', 'zhaoliu@example.com', '13800000006', 'user', 1),
('user5', 'e10adc3949ba59abbe56e057f20f883e', '钱七', 'https://avatars.githubusercontent.com/u/7?v=4', 'qianqi@example.com', '13800000007', 'user', 1);

-- 插入一些特殊状态的用户
INSERT INTO users (username, password, nickname, avatar, email, mobile, role, status) VALUES
('disabled_user', 'e10adc3949ba59abbe56e057f20f883e', '禁用用户', 'https://avatars.githubusercontent.com/u/8?v=4', 'disabled@example.com', '13800000008', 'user', 0),
('test_user', 'e10adc3949ba59abbe56e057f20f883e', '测试用户', 'https://avatars.githubusercontent.com/u/9?v=4', 'test@example.com', '13800000009', 'user', 1);

-- 插入一些已删除的用户
INSERT INTO users (username, password, nickname, avatar, email, mobile, role, status, deleted_at) VALUES
('deleted_user1', 'e10adc3949ba59abbe56e057f20f883e', '已删除用户1', 'https://avatars.githubusercontent.com/u/10?v=4', 'deleted1@example.com', '13800000010', 'user', 1, NOW()),
('deleted_user2', 'e10adc3949ba59abbe56e057f20f883e', '已删除用户2', 'https://avatars.githubusercontent.com/u/11?v=4', 'deleted2@example.com', '13800000011', 'user', 1, NOW()); 