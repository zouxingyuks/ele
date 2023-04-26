package router

import (
	"ele/router/v1"
	"ele/tools"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
)

func Start() {
	r := gin.Default()
	api := r.Group("/api/v1")
	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	v1.LoadApi(api)
	//fmt.Println(tools.Configs.GetString("Router.port"))
	r.Run(":" + tools.Configs.GetString("Router.port"))
}
