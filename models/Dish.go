package models

import (
	"gorm.io/gorm"
	"time"
)

// Dish 菜品
type Dish struct {
	ID          uint `gorm:"primarykey" json:"id,omitempty" form:"id" binding:"required"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Name        string         `gorm:"type:varchar(255);comment:菜品名称" json:"name,omitempty" form:"name" binding:"required"`       // 菜品名称
	Description string         `gorm:"type:text;comment:菜品描述" json:"description,omitempty" form:"description" binding:"required"` // 菜品描述
	Price       float64        `gorm:"type:double;comment:菜品价格" json:"price,omitempty" form:"price" binding:"required"`           // 菜品价格
	Picture     string         `gorm:"type:varchar(255);comment:菜品图片" json:"picture,omitempty" form:"picture" binding:"required"` // 菜品图片
	MerchantID  uint           `gorm:"comment:所属餐厅id" json:"merchantID,omitempty" form:"merchantID" binding:"required"`           // 所属餐厅id
	Comments    []Comment      // 一个菜品有很多评价，使用外键关联
}
