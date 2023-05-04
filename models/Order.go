package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	Dishes []Dish
	UserID uint `gorm:"comment:下单用户ID、" json:"userID,omitempty" form:"userID" binding:"required"` // 所属餐厅id

}
