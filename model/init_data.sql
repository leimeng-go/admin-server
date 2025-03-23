-- 插入默认管理员账号，密码为 admin123
INSERT INTO users (username, password, nickname, role, status) VALUES 
('admin', '$2a$10$RD6c0YT0TiXMgf/2qAFXpehwuTVoQbLLBa.DCUvIKqFj.Ld7SL9Hy', '管理员', 'admin', 1)
ON DUPLICATE KEY UPDATE id=id; 