package dao

import (
	"ele/tools"
	"errors"
	"gorm.io/gorm"
)

// 模糊搜索，传入查询字段，以及结构值
func FuzzyMatch[T any](key string, values *T) {
	tools.DB.Where("name LIKE ?", "%"+key+"%").Find(&values)
}

// 精准搜索，传入查询结构体
func PerfectMatch[T any](key *T, value *T) error {
	err := tools.DB.First(&value, key).Error
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
func List[T any](values *T) {
	tools.DB.Find(&values)
}
