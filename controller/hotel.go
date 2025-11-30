package controller

import (
	"hotel-server/repository"
	"hotel-server/util"

	"github.com/gin-gonic/gin"
)

// HotelController 酒店控制器
type HotelController struct {
	repo *repository.HotelRepository
}

// NewHotelController 初始化酒店控制器
func NewHotelController() *HotelController {
	return &HotelController{
		repo: &repository.HotelRepository{},
	}
}

// GetHotelList 获取酒店列表API
// @Summary 获取酒店列表
// @Router /api/hotel/list [get]
func (c *HotelController) GetHotelList(ctx *gin.Context) {
	hotelList, err := c.repo.GetHotelList(ctx)
	if err != nil {
		util.Fail(ctx, 500, "获取酒店列表失败："+err.Error())
		return
	}
	// 返回统一格式（与前端期望的{list: [...]}对齐）
	util.Success(ctx, gin.H{"list": hotelList}, "获取酒店列表成功")
}
