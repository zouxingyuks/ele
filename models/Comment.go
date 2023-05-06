package models

import (
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
