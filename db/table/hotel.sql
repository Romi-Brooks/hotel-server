CREATE TABLE hotel (
                       hotel_id SERIAL PRIMARY KEY,          -- PK：对应原“酒店ID”，匹配sys_user的自增ID风格
                       city VARCHAR(20),                     -- 对应原“所在城市”
                       room_count INTEGER,                   -- 对应原“客房数量”
                       hotel_name VARCHAR(50),               -- 对应原“酒店名称”
                       contact_phone CHAR(11),               -- 对应原“联系电话”
                       hotel_address VARCHAR(100)            -- 对应原“酒店地址”
);