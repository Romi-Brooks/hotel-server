package repository

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"hotel-server/config"
	"hotel-server/model"

	"github.com/jackc/pgx/v5" // 确保导入pgx
	"github.com/jackc/pgx/v5/pgxpool"
)

// 全局初始化rand（用于生成预订编号）
func init() {
	rand.Seed(time.Now().UnixNano())
}

// BookingRepository 预订数据访问层
type BookingRepository struct {
	db *pgxpool.Pool // 用pgx的连接池（和config保持一致）
}

// NewBookingRepository 初始化预订仓库（依赖注入DB）
func NewBookingRepository() *BookingRepository {
	return &BookingRepository{
		db: config.DB, // 假设config.DB是pgxpool.Pool实例
	}
}

// GetBookingList 联表查询：预订+客户+房间
func (r *BookingRepository) GetBookingList(ctx context.Context) ([]model.BookingVO, error) {
	var bookings []model.BookingVO
	sql := `
		SELECT 
			b.booking_id, b.customer_id, b.room_number, b.booking_no,
			b.check_in_time, b.check_out_time, b.status, b.create_time, b.update_time,
			c.name AS customer_name, c.phone AS customer_phone
		FROM booking b
		JOIN customer c ON b.customer_id = c.customer_id
		JOIN room r ON b.room_number = r.room_number
		ORDER BY b.create_time DESC
	`

	// PGX用Query而非QueryContext
	rows, err := r.db.Query(ctx, sql)
	if err != nil {
		return nil, fmt.Errorf("查询预订列表失败: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var vo model.BookingVO
		err := rows.Scan(
			&vo.BookingId, &vo.CustomerId, &vo.RoomNumber, &vo.BookingNo,
			&vo.CheckInTime, &vo.CheckOutTime, &vo.Status, &vo.CreateTime, &vo.UpdateTime,
			&vo.CustomerName, &vo.CustomerPhone,
		)
		if err != nil {
			return nil, fmt.Errorf("扫描预订数据失败: %w", err)
		}
		bookings = append(bookings, vo)
	}
	return bookings, nil
}

// AddBookingWithRoomStatus 事务：新增预订 + 更新房间状态（PGX事务适配）
func (r *BookingRepository) AddBookingWithRoomStatus(ctx context.Context, booking model.Booking) error {
	// PGX开启事务（用pgx.TxOptions）
	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return fmt.Errorf("开启事务失败: %w", err)
	}
	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback(ctx) // PGX的Rollback需要传ctx
			panic(p)
		} else if err != nil {
			_ = tx.Rollback(ctx)
		} else {
			err = tx.Commit(ctx) // PGX的Commit需要传ctx
		}
	}()

	// 生成预订编号
	bookingNo := fmt.Sprintf("B%s%04d", time.Now().Format("20060102"), rand.Intn(10000))

	// PGX用Exec而非ExecContext
	_, err = tx.Exec(ctx, `
		INSERT INTO booking (customer_id, room_number, booking_no, check_in_time, check_out_time, status)
		VALUES ($1, $2, $3, $4, $5, '待入住')
	`, booking.CustomerId, booking.RoomNumber, bookingNo, booking.CheckInTime, booking.CheckOutTime)
	if err != nil {
		return fmt.Errorf("插入预订记录失败: %w", err)
	}

	// 更新房间状态
	_, err = tx.Exec(ctx, `
		UPDATE room SET current_status = '已预订' WHERE room_number = $1
	`, booking.RoomNumber)
	if err != nil {
		return fmt.Errorf("更新房间状态失败: %w", err)
	}
	return nil
}

// UpdateBookingStatus 更新预订状态
func (r *BookingRepository) UpdateBookingStatus(ctx context.Context, bookingId int, status string) error {
	_, err := r.db.Exec(ctx, `
		UPDATE booking
		SET status = $1, update_time = CURRENT_TIMESTAMP
		WHERE booking_id = $2
	`, status, bookingId)
	if err != nil {
		return fmt.Errorf("更新预订状态失败: %w", err)
	}
	return nil
}

// GetRoomNumberByBookingId 新增：通过预订ID获取房间号（解决controller未解析引用）
func (r *BookingRepository) GetRoomNumberByBookingId(ctx context.Context, bookingId int) (string, error) {
	var roomNumber string
	err := r.db.QueryRow(ctx, `
		SELECT room_number FROM booking WHERE booking_id = $1
	`, bookingId).Scan(&roomNumber)
	if err != nil {
		return "", fmt.Errorf("查询预订对应的房间号失败: %w", err)
	}
	return roomNumber, nil
}
