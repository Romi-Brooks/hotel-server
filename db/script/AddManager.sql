-- 插入管理员（密码：123456，已通过bcrypt加密）
INSERT INTO sys_user (username, password, role)
VALUES ('romi', '$2a$10$Gz7zL6Z9s8e9X5Y7w3Q2R1tK4J6H8N0M7B5V3C1X9Z7S5D3F8G0H', 'admin');