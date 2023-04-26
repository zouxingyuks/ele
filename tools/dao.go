package tools

import (
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var daoLog *logrus.Entry

func initDao() {
	dsn := Configs.GetString("dao.username") + ":" + Configs.GetString("dao.password") + "@tcp(" + Configs.GetString("dao.host") + ")/" + Configs.GetString("dao.dbname") + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		daoLog.Errorln("数据库连接失败")
		panic(err)
	}
	DB = db
}
