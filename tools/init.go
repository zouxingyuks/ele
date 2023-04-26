package tools

// 由于各个功能模块之间存在功能依赖关系
func InitTools() {
	loadDefaultConfig()
	initLog()
	parseConfig()
	initDao()
}
