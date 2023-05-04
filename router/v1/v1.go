package v1

import (
	"ele/controller"
	"github.com/gin-gonic/gin"
)

var api *gin.RouterGroup

func LoadApi(v1 *gin.RouterGroup) {
	api = v1
	// 用户认证授权相关接口
	api.POST("/users", controller.CreateUser)        // 创建用户
	api.POST("/users/login", controller.UserLogin)   // 用户登录
	api.POST("/users/logout", controller.UserLogout) // 用户注销
	loadMerchant()
	loadDish()
	loadOrder()

	// 骑手相关接口
	api.POST("/riders/orders/:order_id", controller.AcceptOrder)            // 骑手接单
	api.POST("/riders/orders/:order_id/complete", controller.CompleteOrder) // 骑手完成订单

	// 评论相关接口
	api.POST("/comments/orders/:order_id", controller.AddComment) // 用户评价订单
}
