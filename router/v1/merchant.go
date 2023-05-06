package v1

import "ele/controller"

// 商家相关 API
func loadMerchant() {
	api.POST("/merchants", controller.AddMerchant)            // 添加商家
	api.PUT("/merchants", controller.UpdateMerchant)          // 商家信息更新
	api.GET("/merchants", controller.ListMerchant)            // 获取所有商家
	api.GET("/merchants/perfect", controller.PerfectMerchant) // 根据名称准确搜索商家详情
	api.GET("/merchants/fuzzy", controller.FuzzyMerchant)     // 根据名称模糊搜索商家详情
	api.DELETE("/merchants", controller.DeleteMerchant)       // DeleteMerchant 根据 id 删除指定商家，及其菜品

}
