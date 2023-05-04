package controller

import "github.com/gin-gonic/gin"

// @Summary 创建订单
// @Description 用户创建新订单
// @Accept json
// @Produce json
// @Param order body string true "新订单信息"
// @Success 200 {object} string "Order"
// @Failure 400 {object} string "ErrorResponse"
// @Failure 500 {object} string "ErrorResponse"
// @Router /orders [post]
func CreateOrder(c *gin.Context) {

}

// @Summary 列出所有订单
// @Description 获取所有订单列表
// @Produce json
// @Success 200 {object} string "succeed"
// @Failure 500 {object} string "ErrorResponse"
// @Router /orders [get]
func ListOrders(c *gin.Context) {

}

// GetOrder 获取订单详情
// @Summary 获取订单详情
// @Description 获取指定id的订单详情
// @Tags 订单管理
// @Accept json
// @Produce json
// @Param order_id path int true "订单id"
// @Success 200 {object} string "OrderDetailResponse"
// @Failure 400 {object} string "ErrorResponse"
// @Router /api/v1/orders/{order_id} [get]
func GetOrder(c *gin.Context) {

}

// CancelOrder 取消订单
// @Summary 取消订单
// @Description 取消指定id的订单
// @Tags 订单管理
// @Accept json
// @Produce json
// @Param order_id path int true "订单id"
// @Success 204 {object} string "succeed"
// @Failure 400 {object} string "ErrorResponse"
// @Router /api/v1/orders/{order_id} [delete]
func CancelOrder(c *gin.Context) {

}
