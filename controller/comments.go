package controller

import "github.com/gin-gonic/gin"

// CreateComment 用户评价订单
// @Summary 用户评价订单
// @Description 用户为指定id的订单添加评价
// @Tags 评价管理
// @Accept json
// @Produce json
// @Param order_id path int true "订单id"
// @Param comment body CreateCommentRequest true "评论信息"
// @Success 201 {object} CommentResponse
// @Failure 400 {object} ErrorResponse
// @Router /api/v1/comments/orders/{order_id} [post]
func CreateComment(c *gin.Context) {

}
