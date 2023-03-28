package controller

import "github.com/gin-gonic/gin"

// CreateUser
// @Summary 创建用户
// @Description 创建新用户
// @Accept json
// @Produce json
// @Param user body User true "新用户信息"
// @Success 200 {object} User
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /users [post]
func CreateUser(c *gin.Context) {

}

// UserLogin 用户登录
// @Summary 用户登录
// @Description 用户使用用户名和密码登录
// @Accept json
// @Produce json
// @Param login body UserLogin true "用户登录信息"
// @Success 200 {object} User
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /users/login [post]
func UserLogin(c *gin.Context) {

}

// UserLogout 用户登出
// @Summary 用户登出
// @Description 用户登出
// @Produce json
// @Success 204 ""
// @Failure 500 {object} ErrorResponse
// @Router /users/logout [post]
func UserLogout(c *gin.Context) {

}
