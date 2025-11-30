package model

// Hotel 对应数据库`hotel`表
type Hotel struct {
	HotelId      int    `json:"hotelId"`      // 酒店ID（主键）
	City         string `json:"city"`         // 城市
	RoomCount    int    `json:"roomCount"`    // 房间总数
	HotelName    string `json:"hotelName"`    // 酒店名称
	ContactPhone string `json:"contactPhone"` // 联系电话
	HotelAddress string `json:"hotelAddress"` // 酒店地址
}
