package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zouxingyuks/tools/check"
	"github.com/zouxingyuks/tools/dao"
	"gorm.io/gorm"
	"time"
)

// Comment 评论
type Comment struct {
	ID         uint `gorm:"primarykey" json:"id,omitempty" form:"id" binding:"required"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	Content    string         `gorm:"type:text;comment:评论内容" json:"content"  form:"content" binding:"required"`        // 评论内容
	Score      int            `gorm:"comment:评论评分" json:"score,omitempty" form:"score" binding:"required"`             // 评论评分
	DishID     uint           `gorm:"comment:评论所属菜品id" json:"dishID,omitempty" form:"dishID" binding:"required"`       // 评论所属菜品id
	CustomerID uint           `gorm:"comment:评论用户ID" json:"customerID,omitempty" form:"customerID" binding:"required"` // 所属餐厅id
	OrderID    uint           `gorm:"comment:对应订单号" json:"orderID,omitempty" form:"orderID" binding:"required"`        // 所属餐厅id

	//images []string 可能有空会多设计个图片的数据？？？
}

func (comment Comment) Update(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (comment Comment) List(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (comment Comment) Perfect(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (comment Comment) Fuzzy(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (comment Comment) Delete(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

// Add 用户评价订单
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
func (comment Comment) Add(c *gin.Context) {
	c.ShouldBind(&comment)
	//中文校验
	if check.CheckChinese(&comment.Content) {
		c.JSON(400, "仅允许中文、英文字母、数字和空白字符，和常见标点符号，不允许输入特殊字符")
		return
	}
	//评分校验
	if !(comment.DishID <= 5 && comment.DishID >= 1) {
		c.JSON(400, "评分只能为 1-5")
		return
	}
	//菜品存在性检验
	var dishes []Dish
	err := dao.PerfectMatch(&Dish{
		ID: comment.DishID,
	}, &dishes)
	if err != nil {
		c.JSON(400, "菜品不存在")
		return
	}
	//todo 是否进行订单存在性校验
	var orders []Order
	err = dao.PerfectMatch(&Order{
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
