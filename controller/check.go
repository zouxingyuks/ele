package controller

import (
	"github.com/gin-gonic/gin"
)

// 查询类数据的检查
func findCheck[T any](c *gin.Context, values []T, err error) {
	//查询失败
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
func addCheck(c *gin.Context, err error) {
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, nil)
}

// 删除操作的结果检查
func delCheck(c *gin.Context, err error) {
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, "删除成功")
}
