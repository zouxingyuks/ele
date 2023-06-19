package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var customers = gin.H{
	"login":    "/api/v1/customers/login",
	"register": "/api/v1/customers",
}

func LoadPages(r *gin.Engine) {
	r.Static("/assets/css", "ui/assets/css")
	r.Static("/assets/images", "ui/assets/images")
	r.LoadHTMLGlob("ui/templates/*.html")

	//登录注册页面
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"api": gin.H{
				"customers": customers,
			}},
		)
	})
	r.GET("/register", func(c *gin.Context) {
		c.HTML(http.StatusOK, "register.html", gin.H{
			"api": gin.H{
				"customers": customers,
			}},
		)
	})
	//拦截器

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	//order pages
	r.GET("/order/list", func(c *gin.Context) {
		c.HTML(http.StatusOK, "order_list.html", nil)
	})
	r.GET("/order/confirm", func(c *gin.Context) {
		c.HTML(http.StatusOK, "order_confirm.html", nil)
	})
	r.GET("/order/pay", func(c *gin.Context) {
		c.HTML(http.StatusOK, "order_pay.html", nil)
	})
	//merchant pages
	r.GET("/merchant/list", func(c *gin.Context) {
		c.HTML(http.StatusOK, "merchant_list.html", nil)
	})
	r.GET("/merchant/detail", func(c *gin.Context) {
		c.HTML(http.StatusOK, "merchant_detail.html", nil)
	})
}
