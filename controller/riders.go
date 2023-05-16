package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

// Rider 骑手
type Rider struct {
	ID        uint `gorm:"primarykey" json:"id,omitempty" form:"id" binding:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string         `gorm:"column:name;type:varchar(255);comment:骑手姓名"`     // 骑手姓名
	Phone     string         `gorm:"column:phone;type:varchar(20);comment:骑手电话"`     // 骑手电话
	Password  string         `gorm:"column:password;type:varchar(255);comment:骑手密码"` // 骑手密码
	//Latitude  float64 `gorm:"column:latitude;type:decimal(10,6);comment:骑手所在纬度"`  // 骑手所在纬度
	//Longitude float64 `gorm:"column:longitude;type:decimal(10,6);comment:骑手所在经度"` // 骑手所在经度
}

func (r Rider) Add(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (r Rider) Update(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (r Rider) List(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (r Rider) Perfect(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (r Rider) Fuzzy(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (r Rider) Delete(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

// AcceptOrder 骑手接单
// @Summary 骑手接单
// @Description 骑手接收指定id的订单
// @Tags 骑手管理
// @Accept json
// @Produce json
// @Param order_id path int true "订单id"
// @Success 204 {object} string "接单成功"
// @Failure 400 {object} string "ErrorResponse"
// @Router /api/v1/riders/orders/{order_id}/accept [post]
func AcceptOrder(c *gin.Context) {

}

// CompleteOrder 骑手完成订单
// @Summary 骑手完成订单
// @Description 骑手完成指定id的订单
// @Tags 骑手管理
// @Accept json
// @Produce json
// @Param order_id path int true "订单id"
// @Success 204 {object} string "订单完成"
// @Failure 400 {object} string "ErrorResponse"
// @Router /api/v1/riders/orders/{order_id}/complete [post]
func CompleteOrder(c *gin.Context) {

}
