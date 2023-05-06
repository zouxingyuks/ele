package v1

import "ele/controller"

func loadComment() {
	// 评论相关接口
	//todo 路由重新设计
	api.POST("/comments", controller.AddComment)     // 用户评价订单
	api.PATCH("/comments", controller.UpdateComment) // 用户更新订单评价

}
