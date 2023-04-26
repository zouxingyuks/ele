package dao

import (
	"ele/tools"
	"errors"
	"gorm.io/gorm"
)

// 模糊搜索，传入查询字段，以及结构值
func FuzzyMatch[T any](key string, values *T, preload ...string) {
	db := tools.DB
	//加载外键
	for _, v := range preload {
		db = db.Preload(v)
	}
	db.Where("name LIKE ?", "%"+key+"%").Find(&values)
}

// 精准搜索，传入查询结构体
func PerfectMatch[T any](key *T, value *T, preload ...string) error {
	db := tools.DB
	//加载外键
	for _, v := range preload {
		db = db.Preload(v)
	}
	err := db.First(&value, key).Error
	//未查询到数据
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	return nil
}
func Add[T any](m *T) error {
	result := tools.DB.Create(&m)
	return result.Error
}

// 列出所有数据
func List[T any](values *T, preload ...string) {
	db := tools.DB
	//加载外键
	for _, v := range preload {
		db = db.Preload(v)
	}
	db.Find(&values)
}
