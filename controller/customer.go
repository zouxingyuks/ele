package controller

import (
	"ele/models"
	"ele/tools"
	"ele/tools/dao"
	"github.com/gin-gonic/gin"
	"strconv"
)

// AddCustomer 添加用户
// @Summary 添加用户
// @Description 添加用户
// @Tags 用户管理
// @Accept multipart/form-data
// @Produce application/json
// @Param username formData string true "用户名"
// @Param password formData string true "密码"
// @Param phone formData string true "手机号"
// @Param address formData string true "用户地址"
// @Success 200 {object} string "添加成功"
// @Success 400 {object} string "输入非法"
// @Success 500 {object} string "添加失败"
// @Router /customers [post]
func AddCustomer(c *gin.Context) {
	var u models.Customer
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

// LoginCustomer 用户登录
// @Summary 用户登录
// @Tags 用户管理
// @Description 用户使用用户名和密码登录
// @Accept multipart/form-data
// @Produce application/json
// @Param phone query string true "手机号"
// @Param password query string true "密码" format(password)
// @Success 200 {object} string "添加成功"
// @Success 400 {object} string "输入非法"
// @Success 401 {object} string "输入非法"
// @Success 500 {object} string "添加失败"
// @Router /customers/login [get]
func LoginCustomer(c *gin.Context) {
	var (
		u        models.Customer
		values   []models.Customer
		password string
	)
	u.Phone = c.Query("phone")
	password = c.Query("password")
	if u.Phone == "" || password == "" {
		c.JSON(400, "输入参数不能为空")
		return
	}
	//手机号有效性检验
	err := tools.CheckPhoneNumber(&u.Phone)
	if err != nil {
		c.JSON(400, u.Phone+err.Error())
		return

	}
	//对应账户查询
	err = dao.PerfectMatch(&u, &values)
	if err != nil && err.Error() != "" {
		c.JSON(500, err)

	} else {
		if len(values) == 0 {
			//资源未找到
			c.JSON(404, "账户不存在")
			return
		}
		passwordCheck(c, password, values[0].Password)

	}
}

// @Summary 用户登出
// @Description 用户登出
// @Tags 用户管理
// @Produce json
// @Success 204 {object} string "退出成功"
// @Failure 500 {object} string "ErrorResponse"
// @Router /customers/logout [post]
func LogoutCustomer(c *gin.Context) {

}

// DeleteCustomer 根据 id 删除指定用户，及其相关信息
// @Tags 用户管理
// @Summary 根据 id 删除指定用户，及其相关信息
// @Description 根据 id 删除指定用户，及其相关信息，包括其订单和收藏的菜品
// @Produce json
// @Param id query int true "用户id"
// @Success 200 {object} string "删除成功"
// @Success 400 {object} string "输入非法"
// @Success 500 {object} string "删除失败"
// @Router /customer [delete]
func DeleteCustomer(c *gin.Context) {

	id, err := strconv.Atoi(c.Query("id"))
	if id == 0 || err != nil {
		c.JSON(400, nil)
		return
	}

	customer := models.Customer{
		ID: uint(id),
	}
	//级联删除用户信息，包括其订单和收藏的菜品
	//主要是为了保护用户隐私
	err = dao.Del(&customer, 3)
	delCheck(c, err)

}
