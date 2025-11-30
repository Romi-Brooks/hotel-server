package repository

import (
	"context"
	"fmt"
	"hotel-server/config"
	"hotel-server/model"
)

// HotelRepository 酒店数据访问层
type HotelRepository struct{}

// GetHotelList 获取所有酒店列表
func (r *HotelRepository) GetHotelList(ctx context.Context) ([]model.Hotel, error) {
	// 查询hotel表所有记录
	sql := `SELECT hotel_id, city, room_count, hotel_name, contact_phone, hotel_address FROM hotel;`

	rows, err := config.DB.Query(ctx, sql)
	if err != nil {
		return nil, fmt.Errorf("查询酒店列表失败: %w", err)
	}
	defer rows.Close()

	var hotelList []model.Hotel
	for rows.Next() {
		var hotel model.Hotel
		// 扫描顺序需与SQL查询字段顺序一致
		err := rows.Scan(
			&hotel.HotelId,
			&hotel.City,
			&hotel.RoomCount,
			&hotel.HotelName,
			&hotel.ContactPhone,
			&hotel.HotelAddress,
		)
		if err != nil {
			return nil, fmt.Errorf("扫描酒店数据失败: %w", err)
		}
		hotelList = append(hotelList, hotel)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("遍历酒店结果失败: %w", err)
	}

	return hotelList, nil
}
