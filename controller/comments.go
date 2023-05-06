package controller

import (
	"ele/models"
	"ele/tools"
	"ele/tools/dao"
	"github.com/gin-gonic/gin"
)

// AddComment 用户评价订单
// @Summary 用户评价订单
// @Description 用户为指定id的订单添加评价
// @Tags 评价管理
// @Accept multipart/form-data
// @Produce application/json
// @Param comment formData string true "评价"
// @Param score formData int true "评分(1-5)"
// @Param dishID formData int  true "对应菜品"
// @Param customerID formData int  true "评论者"
// @Param orderID formData int  true "对应订单ID"
// @Success 200 {object} string "添加成功"
// @Success 400 {object} string "输入非法"
// @Success 500 {object} string "添加失败"
// @Router /comments [post]
func AddComment(c *gin.Context) {
	var comment models.Comment
	c.ShouldBind(&comment)
	//中文校验
	err := tools.CheckChinese(&comment.Content)
	if err != nil {
		c.JSON(400, "菜品名称"+err.Error())
		return
	}
	//评分校验
	if !(comment.DishID <= 5 && comment.DishID >= 1) {
		c.JSON(400, "评分只能为 1-5")
		return
	}
	//菜品存在性检验
	var dishes []models.Dish
	err = dao.PerfectMatch(&models.Dish{
		ID: comment.DishID,
	}, &dishes)
	if err != nil {
		c.JSON(400, "菜品不存在")
		return
	}
	//todo 是否进行订单存在性校验
	var orders []models.Order
	err = dao.PerfectMatch(&models.Order{
		ID: comment.OrderID,
	}, &orders)
	if err != nil {
		c.JSON(400, "订单不存在")
		return
	}

	//所有校验通过
	err = dao.Add(&comment)
	addCheck(c, err)
}

// UpdateComment 更新评价订单
// @Summary 更新评价订单
// @Description 用户为指定id的订单添加评价
// @Tags 评价管理
// @Accept multipart/form-data
// @Produce application/json
// @Param comment formData string true "评价"
// @Param score formData int true "评分"
// @Param dishID formData int  true "所属餐厅"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Router /api/v1/comments/orders/{order_id} [patch]
func UpdateComment(c *gin.Context) {

}
