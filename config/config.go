package config

import (
	"github.com/spf13/viper"
	"strings"
)

func ParseConfig() {
	//设定默认值
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
