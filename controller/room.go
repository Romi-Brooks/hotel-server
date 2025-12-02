package controller

import (
	"hotel-server/repository"
	"hotel-server/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RoomController 房间管理控制器
type RoomController struct {
	repo *repository.RoomRepository
}

func NewRoomController() *RoomController {
	return &RoomController{
		repo: repository.NewRoomRepository(), // 初始化仓库，避免repo为nil
	}
}

// GetRoomList 获取房间列表（联表数据）
func (c *RoomController) GetRoomList(ctx *gin.Context) {
	roomVOList, err := c.repo.GetRoomListWithRelation(ctx)
	if err != nil {
		util.Fail(ctx, 500, "获取房间列表失败："+err.Error())
		return
	}
	util.Success(ctx, gin.H{"list": roomVOList}, "获取房间列表成功")
}

// AddRoom 新增房间（参数匹配实际表字段）
func (c *RoomController) AddRoom(ctx *gin.Context) {
	var req struct {
		RoomNumber    string  `json:"roomNumber" binding:"required"` // 房间号（主键）
		CurrentStatus string  `json:"currentStatus" binding:"required"`
		Price         float64 `json:"price" binding:"required"`
		RoomTypeId    int     `json:"roomTypeId" binding:"required"` // 关联room_type的外键
		HotelId       int     `json:"hotelId" binding:"required"`    // 关联hotel的外键
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		util.Fail(ctx, http.StatusBadRequest, "参数错误："+err.Error())
		return
	}

	err := c.repo.CreateRoom(ctx, struct {
		RoomNumber    string  `json:"roomNumber"`
		CurrentStatus string  `json:"currentStatus"`
		Price         float64 `json:"price"`
		RoomTypeId    int     `json:"roomTypeId"`
		HotelId       int     `json:"hotelId"`
	}(req))
	if err != nil {
		util.Fail(ctx, 500, "新增房间失败："+err.Error())
		return
	}
	util.Success(ctx, gin.H{"room": req}, "新增房间成功")
}

// EditRoom 编辑房间
func (c *RoomController) EditRoom(ctx *gin.Context) {
	var req struct {
		RoomNumber    string  `json:"roomNumber" binding:"required"` // 房间号（主键，用于定位）
		CurrentStatus string  `json:"currentStatus" binding:"required"`
		Price         float64 `json:"price" binding:"required"`
		RoomTypeId    int     `json:"roomTypeId" binding:"required"`
		HotelId       int     `json:"hotelId" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		util.Fail(ctx, http.StatusBadRequest, "参数错误："+err.Error())
		return
	}

	err := c.repo.UpdateRoom(ctx, struct {
		RoomNumber    string  `json:"roomNumber"`
		CurrentStatus string  `json:"currentStatus"`
		Price         float64 `json:"price"`
		RoomTypeId    int     `json:"roomTypeId"`
		HotelId       int     `json:"hotelId"`
	}(req))
	if err != nil {
		util.Fail(ctx, 500, "编辑房间失败："+err.Error())
		return
	}
	util.Success(ctx, gin.H{"room": req}, "编辑房间成功")
}

// DeleteRoom 删除房间（路径参数是room_number）
func (c *RoomController) DeleteRoom(ctx *gin.Context) {
	// 获取路径参数（房间号，string类型）
	roomNumber := ctx.Param("roomNumber")
	if roomNumber == "" {
		util.Fail(ctx, http.StatusBadRequest, "房间号不能为空")
		return
	}

	err := c.repo.DeleteRoom(ctx, roomNumber)
	if err != nil {
		util.Fail(ctx, 500, "删除房间失败："+err.Error())
		return
	}
	util.Success(ctx, gin.H{"roomNumber": roomNumber}, "删除房间成功")
}

func (c *RoomController) GetFreeRoomList(ctx *gin.Context) {
	// 此时c.repo已初始化，不会nil
	freeRooms, err := c.repo.GetFreeRoomList(ctx)
	if err != nil {
		util.Fail(ctx, 500, "获取空闲房间列表失败："+err.Error())
		return
	}
	util.Success(ctx, gin.H{"list": freeRooms}, "获取空闲房间列表成功")
}

// GetRoomDetail 按房间号查询详情（含使用人信息）
func (c *RoomController) GetRoomDetail(ctx *gin.Context) {
	// 获取路径参数：房间号
	roomNumber := ctx.Param("roomNumber")
	if roomNumber == "" {
		util.Fail(ctx, http.StatusBadRequest, "房间号不能为空")
		return
	}

	detail, err := c.repo.GetRoomDetailByNumber(ctx, roomNumber)
	if err != nil {
		util.Fail(ctx, 500, "获取房间详情失败："+err.Error())
		return
	}

	util.Success(ctx, detail, "获取房间详情成功")
}
