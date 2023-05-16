package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zouxingyuks/tools/dao"
	"gorm.io/gorm"
	"strconv"
	"strings"
	"time"
)

type Order struct {
	ID         uint `gorm:"primarykey" json:"id,omitempty" form:"id" binding:"required"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	Dishes     string         `gorm:"type:varchar(255);comment:菜品列表" json:"dishes" form:"dishes" binding:"required,gt=0"`
	CustomerID uint           `gorm:"comment:下单用户ID、" json:"customerID,omitempty" form:"customerID" binding:"required"` // 评论者ID
	Comments   []Comment
}

func (o Order) Update(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (o Order) List(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (o Order) Fuzzy(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

// Add 创建订单
// @Summary 创建订单
// @Tags 订单管理
// @Description 用户创建新订单
// @Accept multipart/form-data
// @Produce application/json
// @Param customerID formData int true "下单用户"
// @Param dishes formData []int true "菜品列表"
// @Success 200 {object} string "添加成功"
// @Success 400 {object} string "输入非法"
// @Success 500 {object} string "添加失败"
// @Router /orders [post]
func (o Order) Add(c *gin.Context) {
	c.ShouldBind(&o)
	//用户存在性校验
	var customers []Customer
	err := dao.PerfectMatch(&Customer{ID: o.CustomerID}, &customers)
	if err != nil {
		c.JSON(400, "用户不存在")
		return
	}

	//菜品存在性检验
	dishesStr := strings.Split(o.Dishes, ",")
	for _, str := range dishesStr {
		id, err := strconv.Atoi(str)
		var values []Dish
		err = dao.PerfectMatch(&Dish{ID: uint(id)}, &values)
		if err != nil || len(values) == 0 {
			c.JSON(400, "菜品 "+strconv.Itoa(id)+"不存在")
			return
		}
	}
	//所有校验通过
	err = dao.Add(&o)
	addCheck(c, err)
}

// Delete 取消订单
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
func (o Order) Delete(c *gin.Context) {

	id, err := strconv.Atoi(c.Query("id"))
	if id == 0 || err != nil {
		c.JSON(400, "输入非法")
		return
	}

	o = Order{ID: uint(id)}
	//把订单标记为完成
	err = dao.Del(&o, 0)
	delCheck(c, err)

}

// Perfect 准确获取订单信息
// @Tags 订单管理
// @Summary 准确获取订单信息
// @Description 根据 用户ID 或 订单ID 准确获取订单信息
// @Produce json
// @Param customerID query uint false "用户ID"
// @Param orderID query uint false "订单ID"
// @Success 200 {array} interface{} "Order"
// @Failure 400 {object} string "请求参数不能为空"
// @Failure 404 {object} string "请求资源不存在"
// @Failure 500 {object} string "查询失败"
// @Router /order/perfect [get]
func (o Order) Perfect(c *gin.Context) {
	customerID, _ := strconv.Atoi(c.Query("customerID"))
	id, _ := strconv.Atoi(c.Query("orderID"))
	o.CustomerID = uint(customerID)
	o.ID = uint(id)

	c.ShouldBind(&o)
	if o.CustomerID == 0 || o.ID == 0 {
		c.JSON(400, "请求参数不能为空")
		return
	}
	var values []Order
	err := dao.PerfectMatch(&o, &values, "Comments")
	findCheck(c, values, err)
}
