package v1

import (
	"ele/controller"
	"github.com/gin-gonic/gin"
)

var api *gin.RouterGroup

func LoadApi(v1 *gin.RouterGroup) {
	api = v1
	loadCustomer()
	loadMerchant()
	loadDish()
	loadOrder()
	loadComment()
	// 骑手相关接口
	api.POST("/riders/orders/:order_id", controller.AcceptOrder)            // 骑手接单
	api.POST("/riders/orders/:order_id/complete", controller.CompleteOrder) // 骑手完成订单

}
