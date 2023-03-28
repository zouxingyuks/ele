package controller

import "github.com/gin-gonic/gin"

// @Summary 创建订单
// @Description 用户创建新订单
// @Accept json
// @Produce json
// @Param order body Order true "新订单信息"
// @Success 200 {object} Order
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /orders [post]
func CreateOrder(c *gin.Context) {

}

// @Summary 列出所有订单
// @Description 获取所有订单列表
// @Produce json
// @Success 200 {array} Order
// @Failure 500 {object} ErrorResponse
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
// @Success 200 {object} OrderDetailResponse
// @Failure 400 {object} ErrorResponse
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
// @Success 204 "No Content"
// @Failure 400 {object} ErrorResponse
// @Router /api/v1/orders/{order_id} [delete]
func CancelOrder(c *gin.Context) {

}
