package controller

import (
	"hotel-server/model"
	"hotel-server/repository"
	"hotel-server/util"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// BookingController 预订管理控制器
type BookingController struct {
	bookingRepo *repository.BookingRepository
	roomRepo    *repository.RoomRepository // 注入房间仓库
}

// NewBookingController 初始化预订控制器（注入依赖）
func NewBookingController() *BookingController {
	return &BookingController{
		bookingRepo: repository.NewBookingRepository(),
		roomRepo:    repository.NewRoomRepository(), // 初始化房间仓库
	}
}

// GetBookingList 获取预订列表（联表客户）
func (c *BookingController) GetBookingList(ctx *gin.Context) {
	bookingList, err := c.bookingRepo.GetBookingList(ctx)
	if err != nil {
		util.Fail(ctx, 500, "获取预订列表失败："+err.Error())
		return
	}
	util.Success(ctx, gin.H{"list": bookingList}, "获取预订列表成功")
}

// AddBooking 新增预订（事务）
func (c *BookingController) AddBooking(ctx *gin.Context) {
	var req struct {
		CustomerId   string `json:"customerId" binding:"required"`
		RoomNumber   string `json:"roomNumber" binding:"required"`
		CheckInTime  string `json:"checkInTime" binding:"required"`
		CheckOutTime string `json:"checkOutTime" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		util.Fail(ctx, http.StatusBadRequest, "参数错误："+err.Error())
		return
	}

	// 解析时间
	//checkInTime, err := time.Parse("2006-01-02", req.CheckInTime)
	//if err != nil {
	//	util.Fail(ctx, http.StatusBadRequest, "入住时间格式错误（需为YYYY-MM-DD）")
	//	return
	//}
	checkInTime, err := time.Parse("2006-01-02", req.CheckInTime)
	if err != nil {
		checkInTime, err = time.Parse("2006-01-02T15:04:05", req.CheckInTime)
	}
	checkOutTime, err := time.Parse("2006-01-02", req.CheckInTime)
	if err != nil {
		checkOutTime, err = time.Parse("2006-01-02T15:04:05", req.CheckInTime)
	}
	//checkOutTime, err := time.Parse("2006-01-02", req.CheckOutTime)
	//if err != nil {
	//	util.Fail(ctx, http.StatusBadRequest, "退房时间格式错误（需为YYYY-MM-DD）")
	//	return
	//}
	if checkOutTime.Before(checkInTime) {
		util.Fail(ctx, http.StatusBadRequest, "退房时间不能早于入住时间")
		return
	}

	// 构造预订对象
	booking := model.Booking{
		CustomerId:   req.CustomerId,
		RoomNumber:   req.RoomNumber,
		CheckInTime:  checkInTime,
		CheckOutTime: checkOutTime,
	}

	// 调用事务方法
	err = c.bookingRepo.AddBookingWithRoomStatus(ctx, booking)
	if err != nil {
		util.Fail(ctx, 500, "新增预订失败："+err.Error())
		return
	}
	util.Success(ctx, gin.H{}, "新增预订成功")
}

// UpdateBookingStatus 更新预订状态（同步房间状态）
func (c *BookingController) UpdateBookingStatus(ctx *gin.Context) {
	var req struct {
		ID     int    `json:"id" binding:"required"`
		Status string `json:"status" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		util.Fail(ctx, http.StatusBadRequest, "参数错误："+err.Error())
		return
	}

	// 更新预订状态
	err := c.bookingRepo.UpdateBookingStatus(ctx, req.ID, req.Status)
	if err != nil {
		util.Fail(ctx, 500, "更新预订状态失败："+err.Error())
		return
	}

	// 同步房间状态
	var roomStatus string
	switch req.Status {
	case "已入住":
		roomStatus = "已入住"
	case "已退房":
		roomStatus = "空闲"
	}
	if roomStatus != "" {
		// 获取预订对应的房间号（调用新增的方法）
		roomNumber, err := c.bookingRepo.GetRoomNumberByBookingId(ctx, req.ID)
		if err != nil {
			util.Fail(ctx, 500, "同步房间状态失败："+err.Error())
			return
		}
		// 更新房间状态（调用房间仓库的方法）
		err = c.roomRepo.UpdateRoomStatus(ctx, roomNumber, roomStatus)
		if err != nil {
			util.Fail(ctx, 500, "同步房间状态失败："+err.Error())
			return
		}
	}

	util.Success(ctx, gin.H{"id": req.ID, "status": req.Status}, "更新预订状态成功")
}
