package model

import "time"

// RoomType 对应数据库`room_type`表
type RoomType struct {
	RoomTypeId int       `json:"roomTypeId"` // 房间类型ID（主键）
	TypeName   string    `json:"typeName"`   // 类型名称（如“单人间”）
	BedType    string    `json:"bedType"`    // 床型（如“1.8m大床”）
	Area       float64   `json:"area"`       // 面积
	MaxPeople  int       `json:"maxPeople"`  // 最大入住人数
	Facilities string    `json:"facilities"` // 设施
	TypeDesc   string    `json:"typeDesc"`   // 类型描述
	CreateTime time.Time `json:"createTime"` // 创建时间
}
