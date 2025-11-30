CREATE TABLE payment_record (
                                payment_record_id SERIAL PRIMARY KEY,  -- PK：对应原“支付记录ID”，匹配sys_user的自增ID风格
                                order_number VARCHAR(20) NOT NULL UNIQUE, -- FK,U：对应原“订单编号”，关联订单表
                                payment_method VARCHAR(20),            -- 对应原“支付方式”
                                payment_amount DECIMAL(10,2),          -- 对应原“支付金额”，保留2位小数
                                payment_status VARCHAR(20),            -- 对应原“支付状态”
                                payment_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- 对应原“支付时间”，匹配sys_user的时间默认值
                                transaction_serial_number VARCHAR(30), -- 对应原“流水号”

    -- 外键约束：关联订单表（需确保订单表已命名为"order"，字段为order_number）
                                CONSTRAINT fk_payment_record_order FOREIGN KEY (order_number)
                                    REFERENCES "order"(order_number)
                                    ON DELETE CASCADE
                                    ON UPDATE CASCADE
);