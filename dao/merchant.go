package dao

import (
	"ele/models"
	"errors"
	"gorm.io/gorm"
)

//单个用get
//全部用list

// 查询部分
// NewMerchant  新建商家
func NewMerchant(m *models.Merchant) error {
	result := DB.Create(&m)
	return result.Error
}

// GetMerchant 查询对应名称的商家
func GetMerchant(name string) (m *models.Merchant) {
	err := DB.First(&m, &models.Merchant{Name: name}).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}
	return m
}

// 获取全部商家信息
func ListMerchants() (merchants []models.Merchant) {
	DB.Find(&merchants)
	return merchants
}
