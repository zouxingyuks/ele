package models

import "gorm.io/gorm"

// Rider 骑手
type Rider struct {
	gorm.Model
	Name     string `gorm:"column:name;type:varchar(255);comment:骑手姓名"`     // 骑手姓名
	Phone    string `gorm:"column:phone;type:varchar(20);comment:骑手电话"`     // 骑手电话
	Password string `gorm:"column:password;type:varchar(255);comment:骑手密码"` // 骑手密码
	//Latitude  float64 `gorm:"column:latitude;type:decimal(10,6);comment:骑手所在纬度"`  // 骑手所在纬度
	//Longitude float64 `gorm:"column:longitude;type:decimal(10,6);comment:骑手所在经度"` // 骑手所在经度
}
