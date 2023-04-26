package controller

import "github.com/gin-gonic/gin"

// AcceptOrder 骑手接单
// @Summary 骑手接单
// @Description 骑手接收指定id的订单
// @Tags 骑手管理
// @Accept json
// @Produce json
// @Param order_id path int true "订单id"
// @Success 204 {object} models.Response "接单成功"
// @Failure 400 {object} models.Response "ErrorResponse"
// @Router /api/v1/riders/orders/{order_id}/accept [post]
func AcceptOrder(c *gin.Context) {

}

// CompleteOrder 骑手完成订单
// @Summary 骑手完成订单
// @Description 骑手完成指定id的订单
// @Tags 骑手管理
// @Accept json
// @Produce json
// @Param order_id path int true "订单id"
// @Success 204 {object} models.Response "订单完成"
// @Failure 400 {object} models.Response "ErrorResponse"
// @Router /api/v1/riders/orders/{order_id}/complete [post]
func CompleteOrder(c *gin.Context) {

}
