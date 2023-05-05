package controller

import (
	"ele/models"
	"ele/tools/dao"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

// AddOrder 创建订单
// @Summary 创建订单
// @Tags 订单管理
// @Description 用户创建新订单
// @Accept multipart/form-data
// @Produce application/json
// @Param userID formData int true "下单用户"
// @Param dishes formData []int true "菜品列表"
// @Success 200 {object} string "添加成功"
// @Success 400 {object} string "输入非法"
// @Success 500 {object} string "添加失败"
// @Router /orders/add [post]
func AddOrder(c *gin.Context) {
	var o models.Order
	c.ShouldBind(&o)
	//用户存在性校验
	var users []models.User
	err := dao.PerfectMatch(&models.User{
		Model: gorm.Model{ID: o.UserID},
	}, &users)
	if err != nil {
		c.JSON(400, "用户不存在")
		return
	}

	//菜品存在性检验
	dishesStr := strings.Split(o.Dishes, ",")
	for _, str := range dishesStr {
		id, err := strconv.Atoi(str)
		var values []models.Dish
		err = dao.PerfectMatch(&models.Dish{
			Model: gorm.Model{ID: uint(id)},
		}, &values)
		if err != nil || len(values) == 0 {
			c.JSON(400, "菜品 "+strconv.Itoa(id)+"不存在")
			return
		}
	}
	//所有校验通过
	err = dao.Add(&o)
	addCheck(c, err)
}

// DeleteOrder 取消订单
// @Summary 取消订单
// @Description 取消指定id的订单
// @Tags 订单管理
// @Accept json
// @Produce json
// @Param id query int true "订单id"
// @Success 200 {object} string "删除成功"
// @Success 400 {object} string "输入非法"
// @Success 500 {object} string "删除失败"
// @Router /orders [delete]
func DeleteOrder(c *gin.Context) {

	id, err := strconv.Atoi(c.Query("id"))
	if id == 0 || err != nil {
		c.JSON(400, "输入非法")
		return
	}

	o := models.Order{
		Model: gorm.Model{ID: uint(id)}}
	//把订单标记为完成
	err = dao.Del(&o, 0)
	delCheck(c, err)

}

// PerfectOrder 准确获取订单信息
// @Tags 订单管理
// @Summary 准确获取订单信息
// @Description 根据 用户ID 或 订单ID 准确获取订单信息
// @Produce json
// @Param userID formData uint false "用户ID"
// @Param orderID formData uint false "订单ID"
// @Success 200 {array} interface{} "Order"
// @Failure 400 {object} string "请求参数不能为空"
// @Failure 404 {object} string "请求资源不存在"
// @Failure 500 {object} string "查询失败"
// @Router /order/perfect [post]
func PerfectOrder(c *gin.Context) {
	o := models.Order{}
	c.ShouldBind(&o)
	if o.UserID == 0 || o.ID == 0 {
		c.JSON(400, "请求参数不能为空")
		return
	}
	var values []models.Order
	err := dao.PerfectMatch(&o, &values, "Comments")
	findCheck(c, values, err)
}
