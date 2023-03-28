package controller

import "github.com/gin-gonic/gin"

// @Summary 列出所有菜品
// @Description 获取所有菜品列表
// @Produce json
// @Success 200 {array} Dish
// @Failure 500 {object} ErrorResponse
// @Router /dishes [get]
func ListDishes(c *gin.Context) {

}

// @Summary 获取菜品信息
// @Description 根据 ID 获取菜品信息
// @Produce json
// @Param dish_id path int true "菜品 ID"
// @Success 200 {object} Dish
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /dishes/{dish_id} [get]
func GetDish(c *gin.Context) {

}
