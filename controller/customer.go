package controller

import (
	"hotel-server/repository"
	"hotel-server/util"

	"github.com/gin-gonic/gin"
)

// CustomerController 客户控制器
type CustomerController struct {
	repo *repository.CustomerRepository
}

// NewCustomerController 初始化客户控制器
func NewCustomerController() *CustomerController {
	return &CustomerController{
		repo: &repository.CustomerRepository{},
	}
}

// GetCustomerList 获取客户列表
func (c *CustomerController) GetCustomerList(ctx *gin.Context) {
	list, err := c.repo.GetCustomerList(ctx)
	if err != nil {
		util.Fail(ctx, 500, "获取客户列表失败："+err.Error())
		return
	}
	util.Success(ctx, gin.H{"list": list}, "获取客户列表成功")
}
