package controller

import (
	"hotel-server/repository"
	"hotel-server/util"

	"github.com/gin-gonic/gin"
)

// RoomTypeController 房间类型控制器
type RoomTypeController struct {
	repo *repository.RoomTypeRepository
}

// NewRoomTypeController 初始化房间类型控制器
func NewRoomTypeController() *RoomTypeController {
	return &RoomTypeController{
		repo: &repository.RoomTypeRepository{},
	}
}

// GetRoomTypeList 获取房间类型列表API
// @Summary 获取房间类型列表
// @Router /api/room/type/list [get]
func (c *RoomTypeController) GetRoomTypeList(ctx *gin.Context) {
	roomTypeList, err := c.repo.GetRoomTypeList(ctx)
	if err != nil {
		util.Fail(ctx, 500, "获取房间类型列表失败："+err.Error())
		return
	}
	// 返回统一格式（与前端期望的{list: [...]}对齐）
	util.Success(ctx, gin.H{"list": roomTypeList}, "获取房间类型列表成功")
}
