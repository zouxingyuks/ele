package v1

import "ele/controller"

func loadDish() {
	// 菜品相关接口
	api.POST("/dishes", controller.AddDish)            // 添加商家
	api.GET("/dishes", controller.ListDish)            // 获取菜品列表
	api.GET("/dishes/perfect", controller.PerfectDish) // 根据名称准确搜索菜品详情
	api.GET("/dishes/fuzzy", controller.FuzzyDish)     // 根据名称模糊搜索菜品详情
	api.DELETE("/dishes", controller.DeleteDish)       // 根据名称模糊搜索菜品详情
}
