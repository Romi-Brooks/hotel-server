-- 管理员账户 / 超级管理员账户表
CREATE TABLE sys_user (
                          id SERIAL PRIMARY KEY,                      -- 自增ID
                          username VARCHAR(50) NOT NULL UNIQUE,       -- 账号（唯一）
                          password VARCHAR(100) NOT NULL,             -- 加密后的密码
                          role VARCHAR(20) NOT NULL,                  -- 角色：super_admin/admin
                          create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- 创建时间
                          update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP  -- 更新时间
);
