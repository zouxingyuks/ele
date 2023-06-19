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
func passwordCheck(c *gin.Context, input string, original string) {
	//todo original 解密模式
	if len(input) != len(original) {
		c.JSON(401, "密码错误")
		return
	}
	for k, v := range original {
		if v != rune(input[k]) {
			c.JSON(401, "密码错误")
			return
		}
	}
	//todo 返回token
	//c.JSON(200, "token还没确定用啥算法")
	c.Redirect(302, "/")

	return
}
