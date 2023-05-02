package models

import (
	"gorm.io/gorm"
)

// Merchant 餐厅
type Merchant struct {
	gorm.Model
	Name    string `gorm:"type:varchar(255);comment:餐厅名称" json:"name,omitempty" form:"name" binding:"required"`        // 餐厅名称
	Address string `gorm:"type:varchar(255);comment:餐厅地址" json:"address,omitempty" form:"address" binding:"required" ` // 餐厅地址
	Phone   string `gorm:"type:varchar(20);comment:餐厅电话" json:"phone,omitempty" form:"phone" binding:"required"`       // 餐厅电话
	Dishes  []Dish // 一个餐厅有多个菜品，使用外键关联
}

// Dish 菜品
type Dish struct {
	gorm.Model
	Name        string    `gorm:"type:varchar(255);comment:菜品名称" json:"name,omitempty" form:"name" binding:"required"`       // 菜品名称
	Description string    `gorm:"type:text;comment:菜品描述" json:"description,omitempty" form:"description" binding:"required"` // 菜品描述
	Price       float64   `gorm:"type:double;comment:菜品价格" json:"price,omitempty" form:"price" binding:"required"`           // 菜品价格
	Picture     string    `gorm:"type:varchar(255);comment:菜品图片" json:"picture,omitempty" form:"picture" binding:"required"` // 菜品图片
	MerchantID  uint      `gorm:"comment:所属餐厅id" json:"merchantID,omitempty" form:"merchantID" binding:"required"`           // 所属餐厅id
	Comments    []Comment // 一个菜品有很多评价，使用外键关联
}

// Comment 评论
type Comment struct {
	gorm.Model
	Content string `gorm:"type:text;comment:评论内容"` // 评论内容
	Score   int    `gorm:"comment:评论评分"`           // 评论评分
	DishID  uint   `gorm:"comment:评论所属菜品id"`       // 评论所属菜品id
}

// Rider 骑手
type Rider struct {
	gorm.Model
	Name     string `gorm:"column:name;type:varchar(255);comment:骑手姓名"`     // 骑手姓名
	Phone    string `gorm:"column:phone;type:varchar(20);comment:骑手电话"`     // 骑手电话
	Password string `gorm:"column:password;type:varchar(255);comment:骑手密码"` // 骑手密码
	//Latitude  float64 `gorm:"column:latitude;type:decimal(10,6);comment:骑手所在纬度"`  // 骑手所在纬度
	//Longitude float64 `gorm:"column:longitude;type:decimal(10,6);comment:骑手所在经度"` // 骑手所在经度
}
