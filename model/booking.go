package model

import "time"

// Booking 预订表模型（对应数据库表）
type Booking struct {
	BookingId    int       `json:"bookingId"`    // 预订ID
	CustomerId   string    `json:"customerId"`   // 关联客户表的外键
	RoomNumber   string    `json:"roomNumber"`   // 关联房间表的外键
	BookingNo    string    `json:"bookingNo"`    // 预订编号
	CheckInTime  time.Time `json:"checkInTime"`  // 入住时间
	CheckOutTime time.Time `json:"checkOutTime"` // 退房时间
	Status       string    `json:"status"`       // 预订状态
	CreateTime   time.Time `json:"createTime"`   // 创建时间
	UpdateTime   time.Time `json:"updateTime"`   // 更新时间
}

// BookingVO 预订视图对象（联表后：包含客户信息）
type BookingVO struct {
	Booking              // 嵌入预订表字段
	CustomerName  string `json:"customerName"`  // 来自customer表的姓名
	CustomerPhone string `json:"customerPhone"` // 来自customer表的电话
}
