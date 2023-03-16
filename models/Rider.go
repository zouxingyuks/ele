package models

type Rider struct {
	Id             string `gorm:"id"`
	Name           string `gorm:"name"`
	CompanyProfile string `gorm:"companyprofile"`
}
