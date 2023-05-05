package controller

import (
	"ele/models"
	"ele/tools"
	"ele/tools/dao"
	"github.com/gin-gonic/gin"
)

// AddUser 添加用户
// @Tags 用户管理
// @Summary 添加用户
// @Description 添加用户
// @Accept multipart/form-data
// @Produce application/json
// @Param username formData string true "用户名"
// @Param password formData string true "密码"
// @Param phone formData string true "手机号"
// @Param address formData string true "用户地址"
// @Success 200 {object} string "添加成功"
// @Success 400 {object} string "输入非法"
// @Success 500 {object} string "添加失败"
// @Router /users/add [post]
func AddUser(c *gin.Context) {
	var u models.User
	c.ShouldBind(&u)
	//姓名中文校验
	err := tools.CheckChinese(&u.Username)
	if err != nil {
		c.JSON(400, "用户名"+err.Error())
		return
	}
	//地址中文校验
	err = tools.CheckChinese(&u.Address)
	if err != nil {
		c.JSON(400, "用户地址"+err.Error())
		return
	}
	//手机号码校验
	//手机号码校验
	err = tools.CheckPhoneNumber(&u.Phone)
	if err != nil {
		c.JSON(400, u.Phone+err.Error())
		return

	}
	//todo 密码加密
	//所有校验通过
	err = dao.Add(&u)
	addCheck(c, err)
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
