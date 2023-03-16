package main

import (
	"ele/router"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"

	"strings"
)

// todo swagger文档
func main() {
	parseConfig()
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		// 配置文件发生变更之后会调用的回调函数
		fmt.Println("Config file changed:", e.Name)
	})
	//启动数据库
	//启动路由
	router.Start()
}
func parseConfig() {
	//设定默认值
	viper.SetDefault("username", "admin")
	// 指定配置文件路径
	configPath := "./config/config.yaml"
	temp1 := strings.Split(configPath, "/")
	configName := strings.Split(temp1[len(temp1)-1], ".")
	viper.AddConfigPath(".") // 还可以在工作目录中查找配置
	viper.SetConfigName(configName[0])
	viper.SetConfigType(configName[1])
	if err := viper.ReadInConfig(); err != nil {
		// 配置文件出错
		//todo 设置日志输出
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// 配置文件未找到错误；如果需要可以忽略
			viper.SafeWriteConfigAs(configPath)
			//todo 补充日志创建说明
		} else {
			// 配置文件被找到，但产生了另外的错误
			panic("The configuration file was found, but another error was generated")
		}
	}

	// 配置文件找到并成功解析
}
