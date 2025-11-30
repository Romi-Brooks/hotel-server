package repository

import (
	"context"
	"fmt"
	"hotel-server/config"
)

// DashboardRepository 仪表盘数据访问层
type DashboardRepository struct{}

// GetRoomTotal 从hotel表中求和room_count得到房间总数
func (r *DashboardRepository) GetRoomTotal(ctx context.Context) (int, error) {
	var total int
	// 求和所有酒店的房间数
	sql := `SELECT COALESCE(SUM(room_count), 0) FROM hotel;`
	// COALESCE处理空值，若没有酒店则返回0
	err := config.DB.QueryRow(ctx, sql).Scan(&total)
	if err != nil {
		return 0, fmt.Errorf("查询房间总数失败: %w", err)
	}
	return total, nil
}

// GetRoomFree 从room表中统计current_status为"空闲"的数量
func (r *DashboardRepository) GetRoomFree(ctx context.Context) (int, error) {
	var count int
	sql := `SELECT COALESCE(COUNT(*), 0) FROM room WHERE current_status = '空闲';`
	err := config.DB.QueryRow(ctx, sql).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("查询空闲房间数失败: %w", err)
	}
	return count, nil
}

// GetBookingTotal 从room表中统计current_status为"已预订"的数量（注意：与你写的"已预定"统一为"已预订"）
func (r *DashboardRepository) GetBookingTotal(ctx context.Context) (int, error) {
	var count int
	sql := `SELECT COALESCE(COUNT(*), 0) FROM room WHERE current_status = '已预订';`
	err := config.DB.QueryRow(ctx, sql).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("查询预订房间数失败: %w", err)
	}
	return count, nil
}

// GetCheckInToday 从room表中统计current_status为"已入住"的数量
func (r *DashboardRepository) GetCheckInToday(ctx context.Context) (int, error) {
	var count int
	sql := `SELECT COALESCE(COUNT(*), 0) FROM room WHERE current_status = '已入住';`
	err := config.DB.QueryRow(ctx, sql).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("查询今日入住数失败: %w", err)
	}
	return count, nil
}

// GetRevenueDay 计算日收入：已入住/已预订房间的price求和
func (r *DashboardRepository) GetRevenueDay(ctx context.Context) (float64, error) {
	var revenue float64
	// 求和已入住/已预订房间的价格
	sql := `SELECT COALESCE(SUM(price), 0) FROM room WHERE current_status IN ('已入住', '已预订');`
	err := config.DB.QueryRow(ctx, sql).Scan(&revenue)
	if err != nil {
		return 0, fmt.Errorf("查询日收入失败: %w", err)
	}
	return revenue, nil
}

// GetOccupancyRate 计算入住率：已入住房间数 / 房间表总数量
func (r *DashboardRepository) GetOccupancyRate(ctx context.Context) (float64, error) {
	// 1. 获取已入住房间数
	checkInCount, err := r.GetCheckInToday(ctx)
	if err != nil {
		return 0, err
	}

	// 2. 获取房间表总数量
	var total int
	sql := `SELECT COALESCE(COUNT(*), 0) FROM room;`
	err = config.DB.QueryRow(ctx, sql).Scan(&total)
	if err != nil {
		return 0, fmt.Errorf("查询房间表总数失败: %w", err)
	}

	// 3. 计算入住率（避免除零错误）
	if total == 0 {
		return 0, nil
	}
	occupancyRate := float64(checkInCount) / float64(total)
	return occupancyRate, nil
}
