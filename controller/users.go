package controller

import "github.com/gin-gonic/gin"

// @Summary 创建用户
// @Description 创建新用户
// @Accept json
// @Produce json
// @Param user body string true "新用户信息"
// @Success 200 {object} string "User"
// @Failure 400 {object} string "ErrorResponse"
// @Failure 500 {object} string "ErrorResponse"
// @Router /users [post]
func CreateUser(c *gin.Context) {

}

// @Summary 用户登录
// @Description 用户使用用户名和密码登录
// @Accept json
// @Produce json
// @Param login body string true "用户登录信息"
// @Success 200 {object} string "User"
// @Failure 400 {object} string "ErrorResponse"
// @Failure 500 {object} string "ErrorResponse"
// @Router /users/login [post]
func UserLogin(c *gin.Context) {

}

// @Summary 用户登出
// @Description 用户登出
// @Produce json
// @Success 204 {object} string "退出成功"
// @Failure 500 {object} string "ErrorResponse"
// @Router /users/logout [post]
func UserLogout(c *gin.Context) {

}
