package model

// Room 对应数据库`room`表（实际字段）
type Room struct {
	RoomNumber    string  `json:"roomNumber"`    // 房间号（主键，char(8)）
	CurrentStatus string  `json:"currentStatus"` // 房间状态（varchar(20)）
	Price         float64 `json:"price"`         // 价格（numeric(10,2)）
	RoomTypeId    int     `json:"roomTypeId"`    // 关联`room_type`的外键
	HotelId       int     `json:"hotelId"`       // 关联`hotel`的外键
}

// RoomVO 房间视图对象（包含关联的`room_type`和`hotel`信息）
type RoomVO struct {
	Room                // 嵌入`room`表的字段
	TypeName     string `json:"typeName"`     // 来自`room_type`的类型名
	BedType      string `json:"bedType"`      // 来自`room_type`的床型
	HotelName    string `json:"hotelName"`    // 来自`hotel`的酒店名
	ContactPhone string `json:"contactPhone"` // 来自`hotel`的联系电话
	HotelAddress string `json:"hotelAddress"` // 来自`hotel`的地址
}
