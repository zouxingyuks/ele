package v1

import "ele/controller"

func loadOrder() {
	// 订单相关接口
	api.POST("/orders/add", controller.CreateOrder)       // 创建订单
	api.GET("/orders/list", controller.ListOrders)        // 获取订单列表
	api.GET("/orders/perfect", controller.GetOrder)       // 根据 ID 查询指定订单
	api.DELETE("/orders/perfect", controller.CancelOrder) // 获取用户所有订单
	//某一用户订单
	//已完成订单
	//未完成订单
	//
	//// 菜品相关接口
	//api.POST("/dish/add", controller.AddDish)         // 添加商家
	//api.GET("/dish/list", controller.ListDish)        // 获取菜品列表
	//api.POST("/dish/perfect", controller.PerfectDish) // 根据名称准确搜索菜品详情
	//api.POST("/dish/fuzzy", controller.FuzzyDish)     // 根据名称模糊搜索菜品详情

}
