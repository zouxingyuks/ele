package models

import "gorm.io/gorm"

// User 饿了么用户
type User struct {
	gorm.Model
	Username  string  `gorm:"type:varchar(255);comment:用户名" json:"username,omitempty" form:"username" binding:"required"` // 用户名
	Password  string  `gorm:"type:varchar(255);comment:密码" json:"password,omitempty" form:"password" binding:"required"`  // 密码
	Phone     string  `gorm:"type:varchar(20);comment:手机号;unique" json:"phone,omitempty" form:"phone" binding:"required"` // 手机号
	Address   string  `gorm:"type:varchar(255);comment:用户地址" json:"address,omitempty" form:"address" binding:"required"`  // 用户地址
	Orders    []Order // 一个用户有多个订单，使用外键关联
	Favorites []Dish  `gorm:"many2many:favorite_dishes;"` // 一个用户有多个收藏菜品，使用多对多关联
}
