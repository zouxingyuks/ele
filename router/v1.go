package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
)

func Start() {
	r := gin.Default()
	api := r.Group("/api/v1")
	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	loadApi(api)
	fmt.Println(viper.GetString("Router.port"))
	r.Run(":" + viper.GetString("Router.port"))
}
