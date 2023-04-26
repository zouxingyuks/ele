package v1

import "ele/controller"

// 商家相关 API
func loadMerchant() {
	api.GET("/merchant/list", controller.ListMerchants)      // 获取商家列表
	api.GET("/merchant/perfect", controller.PerfectMerchant) // 根据 ID 获取商家详情
	api.GET("/merchant/fuzzy", controller.FuzzyMerchant)     // 根据 ID 获取商家详情
	api.POST("/merchant/add", controller.AddMerchant)        // 获取商家列表
}
