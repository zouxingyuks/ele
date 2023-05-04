package controller

import "github.com/gin-gonic/gin"

// AddComment 用户评价订单
// @Summary 用户评价订单
// @Description 用户为指定id的订单添加评价
// @Tags 评价管理
// @Accept json
// @Produce json
// @Success 200 {object} string
// @Failure 400 {object} string
// @Router /api/v1/comments/orders/{order_id} [post]
func AddComment(c *gin.Context) {

}
