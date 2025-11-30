package repository

import (
	"context"
	"fmt"
	"hotel-server/config"
	"hotel-server/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

// RoomRepository 房间数据访问层
type RoomRepository struct {
	db *pgxpool.Pool // 与booking仓库统一用pgx连接池
}

func NewRoomRepository() *RoomRepository {
	return &RoomRepository{
		db: config.DB, // 必须确保config.DB已在程序启动时初始化
	}
}

func (r *RoomRepository) GetRoomListWithRelation(ctx context.Context) ([]model.RoomVO, error) {
	// 联表SQL（匹配你的表字段）
	sql := `
		SELECT 
			r.room_number,
			r.current_status,
			r.price,
			r.room_type_id,
			r.hotel_id,
			rt.type_name,
			rt.bed_type,
			h.hotel_name,
			h.contact_phone,
			h.hotel_address
		FROM room r
		INNER JOIN room_type rt ON r.room_type_id = rt.room_type_id
		INNER JOIN hotel h ON r.hotel_id = h.hotel_id
	`

	// 执行查询
	rows, err := config.DB.Query(ctx, sql)
	if err != nil {
		return nil, fmt.Errorf("联表查询房间失败: %w", err)
	}
	defer rows.Close()

	var roomVOList []model.RoomVO
	// 遍历结果集，扫描到RoomVO
	for rows.Next() {
		var vo model.RoomVO
		// 注意：扫描顺序必须和SQL查询的字段顺序完全一致
		err := rows.Scan(
			&vo.RoomNumber,
			&vo.CurrentStatus,
			&vo.Price,
			&vo.RoomTypeId,
			&vo.HotelId,
			&vo.TypeName,
			&vo.BedType,
			&vo.HotelName,
			&vo.ContactPhone,
			&vo.HotelAddress,
		)
		if err != nil {
			return nil, fmt.Errorf("扫描房间数据失败: %w", err)
		}
		roomVOList = append(roomVOList, vo)
	}

	// 检查遍历过程中的错误
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("遍历房间结果失败: %w", err)
	}

	return roomVOList, nil
}

func (r *RoomRepository) CreateRoom(ctx context.Context, req struct {
	RoomNumber    string  `json:"roomNumber"`
	CurrentStatus string  `json:"currentStatus"`
	Price         float64 `json:"price"`
	RoomTypeId    int     `json:"roomTypeId"`
	HotelId       int     `json:"hotelId"`
}) error {
	// 插入SQL（匹配room表字段）
	sql := `
		INSERT INTO room (room_number, current_status, price, room_type_id, hotel_id)
		VALUES ($1, $2, $3, $4, $5)
		ON CONFLICT (room_number) DO NOTHING; -- 避免房间号重复
	`

	// 执行插入
	_, err := config.DB.Exec(ctx, sql,
		req.RoomNumber,
		req.CurrentStatus,
		req.Price,
		req.RoomTypeId,
		req.HotelId,
	)
	if err != nil {
		return fmt.Errorf("新增房间失败: %w", err)
	}
	return nil
}

func (r *RoomRepository) UpdateRoom(ctx context.Context, req struct {
	RoomNumber    string  `json:"roomNumber"`
	CurrentStatus string  `json:"currentStatus"`
	Price         float64 `json:"price"`
	RoomTypeId    int     `json:"roomTypeId"`
	HotelId       int     `json:"hotelId"`
}) error {
	// 更新SQL（根据room_number主键更新）
	sql := `
		UPDATE room
		SET current_status = $1, price = $2, room_type_id = $3, hotel_id = $4
		WHERE room_number = $5;
	`

	// 执行更新
	_, err := config.DB.Exec(ctx, sql,
		req.CurrentStatus,
		req.Price,
		req.RoomTypeId,
		req.HotelId,
		req.RoomNumber,
	)
	if err != nil {
		return fmt.Errorf("编辑房间失败: %w", err)
	}
	return nil
}

func (r *RoomRepository) DeleteRoom(ctx context.Context, roomNumber string) error {
	// 删除SQL（根据room_number主键删除）
	sql := `DELETE FROM room WHERE room_number = $1;`

	// 执行删除
	_, err := config.DB.Exec(ctx, sql, roomNumber)
	if err != nil {
		return fmt.Errorf("删除房间失败: %w", err)
	}
	return nil
}

func (r *RoomRepository) GetFreeRoomList(ctx context.Context) ([]model.RoomVO, error) {
	var freeRooms []model.RoomVO
	sql := `
		SELECT 
			r.room_number, r.current_status, r.price, r.room_type_id, r.hotel_id,
			rt.type_name, rt.bed_type,
			h.hotel_name, h.hotel_address, h.contact_phone
		FROM room r
		JOIN room_type rt ON r.room_type_id = rt.room_type_id
		JOIN hotel h ON r.hotel_id = h.hotel_id
		WHERE r.current_status = '空闲'
	`

	// 此时r.db已初始化，不会nil
	rows, err := r.db.Query(ctx, sql)
	if err != nil {
		return nil, fmt.Errorf("查询空闲房间失败: %w", err)
	}
	defer rows.Close()

	// 扫描结果（后续逻辑不变）
	for rows.Next() {
		var vo model.RoomVO
		err := rows.Scan(
			&vo.RoomNumber, &vo.CurrentStatus, &vo.Price, &vo.RoomTypeId, &vo.HotelId,
			&vo.TypeName, &vo.BedType, &vo.HotelName, &vo.HotelAddress, &vo.ContactPhone,
		)
		if err != nil {
			return nil, fmt.Errorf("扫描空闲房间数据失败: %w", err)
		}
		freeRooms = append(freeRooms, vo)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("遍历空闲房间结果失败: %w", err)
	}
	return freeRooms, nil
}

func (r *RoomRepository) UpdateRoomStatus(ctx context.Context, roomNumber string, status string) error {
	_, err := r.db.Exec(ctx, `
		UPDATE room SET current_status = $1 WHERE room_number = $2
	`, status, roomNumber)
	if err != nil {
		return fmt.Errorf("更新房间状态失败: %w", err)
	}
	return nil
}
