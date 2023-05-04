package v1

import "ele/controller"

func loadDish() {
	// 菜品相关接口
	api.POST("/dish/add", controller.AddDish)         // 添加商家
	api.GET("/dish/list", controller.ListDish)        // 获取菜品列表
	api.POST("/dish/perfect", controller.PerfectDish) // 根据名称准确搜索菜品详情
	api.POST("/dish/fuzzy", controller.FuzzyDish)     // 根据名称模糊搜索菜品详情
	api.DELETE("/dish/", controller.DeleteDish)       // 根据名称模糊搜索菜品详情

}
