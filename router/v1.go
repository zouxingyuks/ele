package router

import (
	_ "ele/docs" // 千万不要忘了导入把你上一步生成的docs
	"fmt"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

func Start() {
	r := gin.Default()
	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	fmt.Println(viper.GetString("Router.port"))
	r.Run(":" + viper.GetString("Router.port"))
}
