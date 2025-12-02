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

	// APIs
	api := r.Group("/api")
	{
		// user
		userCtrl := controller.NewUserController()
		api.POST("/user/login", userCtrl.Login)
		api.POST("/user/addAdmin", userCtrl.AddAdmin)
		api.GET("/user/listAdmin", userCtrl.ListAdmin)

		// room
		roomCtrl := controller.NewRoomController()
		api.GET("/room/list", roomCtrl.GetRoomList)
		api.POST("/room/add", roomCtrl.AddRoom)
		api.PUT("/room/edit", roomCtrl.EditRoom)
		api.DELETE("/room/delete/:roomNumber", roomCtrl.DeleteRoom)
		api.GET("/room/detail/:roomNumber", roomCtrl.GetRoomDetail)
		api.GET("/room/freeList", roomCtrl.GetFreeRoomList)

		// room type
		roomTypeCtrl := controller.NewRoomTypeController()
		api.GET("/room/type/list", roomTypeCtrl.GetRoomTypeList)

		// booking
		bookingCtrl := controller.NewBookingController()
		api.GET("/booking/list", bookingCtrl.GetBookingList)
		api.POST("/booking/updateStatus", bookingCtrl.UpdateBookingStatus)
		api.POST("/booking/add", bookingCtrl.AddBooking)

		// dashboard
		dashboardCtrl := controller.NewDashboardController()
		api.GET("/dashboard", dashboardCtrl.GetDashboardData)

		// hotel
		hotelCtrl := controller.NewHotelController()
		api.GET("/hotel/list", hotelCtrl.GetHotelList)

		// 客户列表路由（确保已注册）
		customerCtrl := controller.NewCustomerController()
		api.GET("/customer/list", customerCtrl.GetCustomerList) // 对应前端的/api/customer/list

	}

	return r
}
