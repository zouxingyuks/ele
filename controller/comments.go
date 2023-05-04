package controller

import (
	"ele/models"
	"ele/tools"
	"ele/tools/dao"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// AddComment 用户评价订单
// @Summary 用户评价订单
// @Description 用户为指定id的订单添加评价
// @Tags 评价管理
// @Accept multipart/form-data
// @Produce application/json
// @Param comment formData string true "评价"
// @Param score formData int true "评分"
// @Param dishID formData int  true "所属餐厅"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Router /api/v1/comments/orders/{order_id} [post]
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
		Model: gorm.Model{ID: comment.DishID},
	}, &dishes)
	if err != nil {
		c.JSON(400, "菜品不存在")
		return
	}

	//所有校验通过
	err = dao.Add(&comment)
	addCheck(c, err)
}
