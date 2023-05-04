package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// 查询类数据的检查
func findCheck[T any](c *gin.Context, values []T, err error) {
	//查询失败
	fmt.Println(err)
	if err != nil && err.Error() != "" {
		c.JSON(500, err)

	} else {
		if len(values) == 0 {
			//资源未找到
			c.JSON(404, nil)
			return
		}
		//返回数据集
		c.JSON(200, values)
	}
}

// 添加操作的检查
