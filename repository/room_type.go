package repository

import (
	"context"
	"fmt"
	"hotel-server/config"
	"hotel-server/model"
)

// RoomTypeRepository 房间类型数据访问层
type RoomTypeRepository struct{}

// GetRoomTypeList 获取所有房间类型列表
func (r *RoomTypeRepository) GetRoomTypeList(ctx context.Context) ([]model.RoomType, error) {
	// 查询room_type表所有记录
	sql := `SELECT room_type_id, type_name, bed_type, area, max_people, facilities, type_desc, create_time FROM room_type;`

	rows, err := config.DB.Query(ctx, sql)
	if err != nil {
		return nil, fmt.Errorf("查询房间类型列表失败: %w", err)
	}
	defer rows.Close()

	var roomTypeList []model.RoomType
	for rows.Next() {
		var roomType model.RoomType
		// 扫描顺序需与SQL查询字段顺序一致
		err := rows.Scan(
			&roomType.RoomTypeId,
			&roomType.TypeName,
			&roomType.BedType,
			&roomType.Area,
			&roomType.MaxPeople,
			&roomType.Facilities,
			&roomType.TypeDesc,
			&roomType.CreateTime,
		)
		if err != nil {
			return nil, fmt.Errorf("扫描房间类型数据失败: %w", err)
		}
		roomTypeList = append(roomTypeList, roomType)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("遍历房间类型结果失败: %w", err)
	}

	return roomTypeList, nil
}
