package main

import (
	"ele/dao"
	"ele/models"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Bind() {
	dao.DB.AutoMigrate(models.MerchantInfo{})
}
func main() {
	parseConfig()
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		// 配置文件发生变更之后会调用的回调函数
		fmt.Println("Config file changed:", e.Name)
	})
	//连接数据库
	dao.Start()
	//模型绑定
	Bind()
}
