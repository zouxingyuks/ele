package controller

import "github.com/gin-gonic/gin"

type CRUD interface {
	Add(c *gin.Context)
	Update(c *gin.Context)
	List(c *gin.Context)
	Find
	Delete(c *gin.Context)
}
type Find interface {
	Perfect(c *gin.Context)
	Fuzzy(c *gin.Context)
}

// NewController 新建控制器
// 拓展时在控制器中新建判断即可
func NewController(s string) CRUD {
	var crud CRUD
	switch s {
	case "customers":
		crud = Customer{}
	case "dishes":
		crud = Dish{}
	case "merchants":
		crud = Merchant{}
	case "orders":
		crud = Order{}
	case "comments":
		crud = Comment{}
	case "riders":
		crud = Rider{}
	default:
		crud = nil
	}
	return crud
}
