package router

import (
	"ele/controller"
	"github.com/gin-gonic/gin"
)

func loadApi(api *gin.RouterGroup) {
	// 用户认证授权相关接口
	api.POST("/users", controller.CreateUser)        // 创建用户
	api.POST("/users/login", controller.UserLogin)   // 用户登录
	api.POST("/users/logout", controller.UserLogout) // 用户注销

	// 商家相关接口
	api.GET("/merchants", controller.ListMerchants)            // 获取商家列表
	api.GET("/merchants/:merchant_id", controller.GetMerchant) // 根据 ID 获取商家详情

	// 菜品相关接口
	api.GET("/dishes", controller.ListDishes)       // 获取菜品列表
	api.GET("/dishes/:dish_id", controller.GetDish) // 根据 ID 获取菜品详情

	// 订单相关接口
	api.POST("/orders", controller.CreateOrder)             // 创建订单
	api.GET("/orders", controller.ListOrders)               // 获取订单列表
	api.GET("/orders/:order_id", controller.GetOrder)       // 根据 ID 获取订单详情
	api.DELETE("/orders/:order_id", controller.CancelOrder) // 取消订单

	// 骑手相关接口
	api.POST("/riders/orders/:order_id", controller.AcceptOrder)            // 骑手接单
	api.POST("/riders/orders/:order_id/complete", controller.CompleteOrder) // 骑手完成订单

	// 评论相关接口
	api.POST("/comments/orders/:order_id", controller.CreateComment) // 用户评价订单
}
