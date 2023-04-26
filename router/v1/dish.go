package v1

import "ele/controller"

func loadDish() {
	// 菜品相关接口
	api.GET("/dish/list", controller.ListDish)      // 获取菜品列表
	api.GET("/dishes/:dish_id", controller.GetDish) // 根据 ID 获取菜品详情

}
