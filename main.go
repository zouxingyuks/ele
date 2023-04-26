package main

import (
	_ "ele/docs" // 千万不要忘了导入把你上一步生成的docs
	"ele/router"
	"ele/tools"
	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
)

// @title 饿了么项目复刻
// @version 1.0
// @description 在东软的教学包基础上去做拓展
// @termsOfService http://swagger.io/terms/

// @contact.name 跑路了
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:9090
// @BasePath /api/v1
var configLog *logrus.Entry

func init() {
	tools.InitTools()
	configLog = tools.NewLog("config")
}
func main() {
	tools.Configs.WatchConfig()
	tools.Configs.OnConfigChange(func(e fsnotify.Event) {
		// 配置文件发生变更之后会调用的回调函数
		configLog.Infoln("Config file changed:", e.Name)
	})
	//启动路由
	router.Start()
}
