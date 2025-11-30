-- 1. 创建房型表
CREATE TABLE IF NOT EXISTS room_type (
                                         room_type_id SERIAL PRIMARY KEY,              -- 自增主键（PostgreSQL专用）
                                         type_name VARCHAR(30) NOT NULL UNIQUE,
                                         bed_type VARCHAR(20) NOT NULL,
                                         area DECIMAL(5,1) NOT NULL,
                                         max_people SMALLINT NOT NULL,                 -- 替代TINYINT，PostgreSQL用SMALLINT
                                         facilities VARCHAR(200) DEFAULT NULL,
                                         type_desc VARCHAR(200) DEFAULT NULL,
                                         create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP -- 时间字段
);