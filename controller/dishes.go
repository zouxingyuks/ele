package controller

import (
	"ele/models"
	"ele/tools/dao"
	"github.com/gin-gonic/gin"
)

// ListDish 列出所有菜品
// @Tags 订单管理
// @Summary 列出所有菜品
// @Description 获取所有菜品列表
// @Produce json
// @Success 200 {object} models.Response "获取成功"
// @Failure 500 {object} models.Response "ErrorResponse"
// @Router /dish/list [get]
func ListDish(c *gin.Context) {
	var values []models.Dish
	dao.List(&values)
	c.JSON(200, models.Response{
		Msg:  "下面是所有菜品信息",
		Data: values,
	})
}

// GetDish 获取菜品信息
// @Summary 获取菜品信息
// @Tags 订单管理
// @Description 根据 ID 获取菜品信息
// @Produce json
// @Param dish_id path int true "菜品 ID"
// @Success 200 {object} models.Response "Dish"
// @Failure 400 {object} models.Response "ErrorResponse"
// @Failure 500 {object} models.Response "ErrorResponse"
// @Router /dishes/{dish_id} [get]
func GetDish(c *gin.Context) {

}
