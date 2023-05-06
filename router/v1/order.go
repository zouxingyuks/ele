package v1

import "ele/controller"

func loadOrder() {
	// 订单相关接口
	api.POST("/orders", controller.AddOrder)      // 创建订单
	api.DELETE("/orders", controller.DeleteOrder) // 取消订单
	//某一用户订单
	//已完成订单
	//未完成订单

}
