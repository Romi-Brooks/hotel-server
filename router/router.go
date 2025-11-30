package router

import (
	"hotel-server/controller" // 导入控制器

	"github.com/gin-gonic/gin"
)

// InitRouter 初始化路由
func InitRouter() *gin.Engine {
	r := gin.Default()

	// 跨域中间件（前端本地开发需跨域）
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// API分组
	api := r.Group("/api")
	{
		// 1. 用户接口
		userCtrl := controller.NewUserController()
		api.POST("/user/login", userCtrl.Login)
		api.POST("/user/addAdmin", userCtrl.AddAdmin)
		api.GET("/user/listAdmin", userCtrl.ListAdmin)

		// 房间接口
		roomCtrl := controller.NewRoomController()
		api.GET("/room/list", roomCtrl.GetRoomList)
		api.POST("/room/add", roomCtrl.AddRoom)
		api.PUT("/room/edit", roomCtrl.EditRoom)
		api.DELETE("/room/delete/:roomNumber", roomCtrl.DeleteRoom) // 路径参数改为roomNumber
		api.GET("/room/freeList", roomCtrl.GetFreeRoomList)         // 对应前端的/api/room/freeList

		// 新增：房间类型列表API
		roomTypeCtrl := controller.NewRoomTypeController()
		api.GET("/room/type/list", roomTypeCtrl.GetRoomTypeList)

		// 3. 预订接口（新增）
		bookingCtrl := controller.NewBookingController()
		api.GET("/booking/list", bookingCtrl.GetBookingList)               // 获取预订列表
		api.POST("/booking/updateStatus", bookingCtrl.UpdateBookingStatus) // 更新预订状态
		api.POST("/booking/add", bookingCtrl.AddBooking)                   // 新增这一行，注册POST接口
		
		dashboardCtrl := controller.NewDashboardController()
		api.GET("/dashboard", dashboardCtrl.GetDashboardData)

		// 新增：酒店列表API
		hotelCtrl := controller.NewHotelController()
		api.GET("/hotel/list", hotelCtrl.GetHotelList)

		// 客户列表路由（确保已注册）
		customerCtrl := controller.NewCustomerController()
		api.GET("/customer/list", customerCtrl.GetCustomerList) // 对应前端的/api/customer/list

	}

	return r
}
