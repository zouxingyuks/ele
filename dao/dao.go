package dao

import (
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func Start() {
	dsn := viper.GetString("Dao.username") + ":" + viper.GetString("Dao.password") + "@tcp(" + viper.GetString("Dao.host") + ")/" + viper.GetString("Dao.dbname") + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("数据库连接失败")
		panic(err)
	}
	DB = db
}
