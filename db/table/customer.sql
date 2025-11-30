-- 客户表
CREATE TABLE customer (
                          customer_id CHAR(18) PRIMARY KEY,                  -- 客户ID（主键）
                          name VARCHAR(20) NOT NULL,                                  -- 姓名
                          phone CHAR(11) UNIQUE,                             -- 手机号（唯一约束）
                          id_card_or_passport CHAR(20) UNIQUE,               -- 身份证号或护照号（唯一约束）
                          email VARCHAR(50) NOT NULL UNIQUE,                 -- 邮箱（非空+唯一约束）
                          create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,   -- 创建时间
                          update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP    -- 更新时间
);