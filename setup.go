package main

import (
	"ele/models"
	"ele/tools"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Bind() {
	_ = tools.DB.AutoMigrate(models.Merchant{})
	_ = tools.DB.AutoMigrate(models.Product{})
	_ = tools.DB.AutoMigrate(models.Comment{})
	_ = tools.DB.AutoMigrate(models.Rider{})
}
func main() {
	tools.InitTools()
	tools.Configs.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		// 配置文件发生变更之后会调用的回调函数
		fmt.Println("Config file changed:", e.Name)
	})
	//模型绑定
	Bind()
}
