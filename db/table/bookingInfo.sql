CREATE TABLE booking (
                         booking_id SERIAL PRIMARY KEY,         -- 预订ID（主键）
                         customer_id CHAR(18) NOT NULL REFERENCES customer(customer_id), -- 关联客户表
                         room_number CHAR(8) NOT NULL REFERENCES room(room_number),       -- 关联房间表
                         booking_no VARCHAR(20) UNIQUE NOT NULL,-- 预订编号（如B20251201001）
                         check_in_time DATE NOT NULL,           -- 入住时间
                         check_out_time DATE NOT NULL,          -- 退房时间
                         status VARCHAR(20) DEFAULT '待入住',   -- 状态：待入住/已入住/已退房/已取消
                         create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                         update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);