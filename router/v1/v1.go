package v1

import (
	"ele/controller"
	"github.com/gin-gonic/gin"
)

var api *gin.RouterGroup

func LoadApi(v1 *gin.RouterGroup) {
	api = v1
	loadCustomer()
	loadCRUD("merchants")
	loadCRUD("dishes")
	loadCRUD("orders")
	loadCRUD("comments")
	loadCRUD("riders")
}
func loadCRUD(s string) {
	crud := controller.NewController(s)
	api.POST("/"+s, crud.Add)
	api.PUT("/"+s, crud.Update)
	api.GET("/"+s, crud.List)
	api.GET("/"+s+"/perfect", crud.Perfect)
	api.GET("/"+s+"/fuzzy", crud.Fuzzy)
	api.DELETE("/"+s, crud.Delete)

}

// 客户类API
func loadCustomer() {
	loadCRUD("customers")
	api.GET("/customers/login", controller.LoginCustomer)    // 用户登录
	api.POST("/customers/logout", controller.LogoutCustomer) // 用户注销
}
