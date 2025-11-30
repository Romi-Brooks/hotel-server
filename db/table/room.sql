CREATE TABLE room (
                      room_number CHAR(8) PRIMARY KEY,                 -- PK：房间号
                      room_type VARCHAR(30) DEFAULT NULL,              -- 房型名称（加长长度）
                      current_status VARCHAR(20) DEFAULT '空闲',       -- 当前状态（默认值：空闲）
                      price DECIMAL(10,2) NOT NULL,                    -- 价格：精确到分，10位总长度，2位小数
                      room_type_id INTEGER DEFAULT NULL,               -- FK：关联房型表（修正命名）
                      hotel_id INTEGER DEFAULT NULL,                   -- FK：关联酒店表（显式允许为空）

    -- 外键约束1：关联房型表
                      CONSTRAINT fk_room_type FOREIGN KEY (room_type_id)
                          REFERENCES room_type(room_type_id)
                          ON DELETE SET NULL
                          ON UPDATE CASCADE,

    -- 外键约束2：关联酒店表
                      CONSTRAINT fk_room_hotel FOREIGN KEY (hotel_id)
                          REFERENCES hotel(hotel_id)
                          ON DELETE SET NULL
                          ON UPDATE CASCADE
);