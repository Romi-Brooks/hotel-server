package controller

import (
	"hotel-server/repository"
	"hotel-server/util"
	"math"

	"github.com/gin-gonic/gin"
)

// DashboardController 仪表盘控制器
type DashboardController struct {
	repo *repository.DashboardRepository // 引入数据访问层
}

// NewDashboardController 初始化仪表盘控制器（创建repository实例）
func NewDashboardController() *DashboardController {
	return &DashboardController{
		repo: &repository.DashboardRepository{},
	}
}

// GetDashboardData 获取仪表盘数据
// @Summary 获取仪表盘数据
// @Router /api/dashboard [get]
func (c *DashboardController) GetDashboardData(ctx *gin.Context) {
	// 1. 从数据库查询各统计指标
	roomTotal, err := c.repo.GetRoomTotal(ctx)
	if err != nil {
		util.Fail(ctx, 500, "获取房间总数失败："+err.Error())
		return
	}

	roomFree, err := c.repo.GetRoomFree(ctx)
	if err != nil {
		util.Fail(ctx, 500, "获取空闲房间数失败："+err.Error())
		return
	}

	bookingTotal, err := c.repo.GetBookingTotal(ctx)
	if err != nil {
		util.Fail(ctx, 500, "获取预订总数失败："+err.Error())
		return
	}

	checkInToday, err := c.repo.GetCheckInToday(ctx)
	if err != nil {
		util.Fail(ctx, 500, "获取今日入住数失败："+err.Error())
		return
	}

	revenueDay, err := c.repo.GetRevenueDay(ctx)
	if err != nil {
		util.Fail(ctx, 500, "获取日收入失败："+err.Error())
		return
	}

	occupancyRate, err := c.repo.GetOccupancyRate(ctx)
	if err != nil {
		util.Fail(ctx, 500, "获取入住率失败："+err.Error())
		return
	}

	// 2. 组装数据（入住率保留2位小数，提升可读性）
	dashboardData := gin.H{
		"roomTotal":     roomTotal,                           // 酒店表room_count求和
		"roomFree":      roomFree,                            // 空闲房间数
		"bookingTotal":  bookingTotal,                        // 已预订房间数
		"checkInToday":  checkInToday,                        // 已入住房间数
		"revenueDay":    math.Round(revenueDay*100) / 100,    // 日收入（保留2位小数）
		"occupancyRate": math.Round(occupancyRate*100) / 100, // 入住率（保留2位小数）
	}

	// 3. 返回成功响应
	util.Success(ctx, dashboardData, "获取仪表盘数据成功")
}
