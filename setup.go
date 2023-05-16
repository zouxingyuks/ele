package main

import (
	"ele/controller"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"github.com/zouxingyuks/tools"
	"github.com/zouxingyuks/tools/config"
	"github.com/zouxingyuks/tools/dao"
)

func Bind() {
	_ = dao.DB.AutoMigrate(controller.Merchant{})
	_ = dao.DB.AutoMigrate(controller.Dish{})
	_ = dao.DB.AutoMigrate(controller.Comment{})
	_ = dao.DB.AutoMigrate(controller.Order{})
	_ = dao.DB.AutoMigrate(controller.Rider{})
	_ = dao.DB.AutoMigrate(controller.Customer{})
}
func main() {
	tools.InitTools()
	config.Configs.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		// 配置文件发生变更之后会调用的回调函数
		fmt.Println("Config file changed:", e.Name)
	})
	//模型绑定
	Bind()
}
