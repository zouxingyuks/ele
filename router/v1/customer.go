package v1

import "ele/controller"

func loadCustomer() {
	api.POST("/customers", controller.AddCustomer)           // 创建用户
	api.GET("/customers/login", controller.LoginCustomer)    // 用户登录
	api.POST("/customers/logout", controller.LogoutCustomer) // 用户注销
	api.DELETE("/customers")
}
