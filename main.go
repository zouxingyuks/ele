package main

import (
	"ele/config"
	"ele/dao"
	"ele/router"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"

	_ "ele/docs" // 千万不要忘了导入把你上一步生成的docs
)

// @title 饿了么项目复刻
// @version 1.0
// @description 在东软的教学包基础上去做拓展
// @termsOfService http://swagger.io/terms/

// @contact.name 这里写联系人信息
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 这里写接口服务的host
// @BasePath 这里写base path

func main() {
	config.ParseConfig()
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		// 配置文件发生变更之后会调用的回调函数
		fmt.Println("Config file changed:", e.Name)
	})
	//启动数据库
	dao.Start()
	//启动路由
	router.Start()
}
