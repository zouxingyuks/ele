package models

// Merchant 商家
type Merchant struct {
	MerchantInfo
	Evaluation
}

// MerchantInfo 基本信息
type MerchantInfo struct {
	Id      string `gorm:"id"`
	Address string `gorm:"address"`
	Name    string `gorm:"name"`
	Imgurl  string `gorm:"imgurl"`
}
type Evaluation struct {
	Stars int `gorm:"stars"`
}
