package tools

// InitTools 由于各个功能模块之间存在功能依赖关系,因此需要手动进行初始化
func InitTools() {
	loadDefaultConfig()
	initLog()
	parseConfig()
	initDao()
}
