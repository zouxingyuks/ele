package models

import (
	"time"
)

// Merchant 餐厅
type Merchant struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt struct {
		Time  time.Time
		Valid bool // Valid is true if Time is not NULL
	} `gorm:"index"`
	Name    string `gorm:"type:varchar(255);comment:餐厅名称" json:"name,omitempty" form:"name" binding:"required"`        // 餐厅名称
	Address string `gorm:"type:varchar(255);comment:餐厅地址" json:"address,omitempty" form:"address" binding:"required" ` // 餐厅地址
	Phone   string `gorm:"type:varchar(20);comment:餐厅电话" json:"phone,omitempty" form:"phone" binding:"required"`       // 餐厅电话
	Dishes  []Dish // 一个餐厅有多个菜品，使用外键关联
}
