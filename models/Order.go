package models

import (
	"gorm.io/gorm"
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
