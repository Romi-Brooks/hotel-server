CREATE TABLE "order" (
                         order_number VARCHAR(20) PRIMARY KEY,  -- PK：对应原“订单编号”
                         reserved_room_count INTEGER,          -- 对应原“预订房间数”
                         guest_count INTEGER,                  -- 对应原“入住人数”
                         order_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- 对应原“订单时间”，匹配sys_user的时间默认值
                         expected_checkin_time TIMESTAMP,      -- 对应原“预计入住时间”
                         expected_checkout_time TIMESTAMP,     -- 对应原“预计退房时间”
                         customer_id CHAR(18) NOT NULL UNIQUE, -- FK,U：对应原“客户ID”，关联客户表
                         hotel_id INTEGER NOT NULL UNIQUE,     -- FK,U：对应原“酒店ID”，关联酒店表

    -- 外键约束1：关联客户表（需确保客户表已命名为customer，字段为customer_id）
                         CONSTRAINT fk_order_customer FOREIGN KEY (customer_id)
                             REFERENCES customer(customer_id)
                             ON DELETE CASCADE
                             ON UPDATE CASCADE,

    -- 外键约束2：关联酒店表（需确保酒店表已命名为hotel，字段为hotel_id）
                         CONSTRAINT fk_order_hotel FOREIGN KEY (hotel_id)
                             REFERENCES hotel(hotel_id)
                             ON DELETE CASCADE
                             ON UPDATE CASCADE
);