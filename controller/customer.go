package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zouxingyuks/tools/check"
	"github.com/zouxingyuks/tools/dao"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type Customer struct {
	ID        uint `gorm:"primarykey" json:"id,omitempty" form:"id" binding:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Username  string         `gorm:"type:varchar(255);comment:用户名" json:"username,omitempty" form:"username" binding:"required"` // 用户名
	Password  string         `gorm:"type:varchar(255);comment:密码" json:"password,omitempty" form:"password" binding:"required"`  // 密码
	Phone     string         `gorm:"type:varchar(20);comment:手机号;unique" json:"phone,omitempty" form:"phone" binding:"required"` // 手机号
	Address   string         `gorm:"type:varchar(255);comment:用户地址" json:"address,omitempty" form:"address" binding:"required"`  // 用户地址
	Orders    []Order        // 一个用户有多个订单，使用外键关联
	Favorites []Dish         `gorm:"many2many:favorite_dishes;"` // 一个用户有多个收藏菜品，使用多对多关联
}

func (customer Customer) Update(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (customer Customer) List(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (customer Customer) Perfect(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (customer Customer) Fuzzy(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

// Add 添加用户
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
func (customer Customer) Add(c *gin.Context) {
	c.ShouldBind(&customer)
	//姓名中文校验
	if check.CheckChinese(&customer.Username) {
		c.JSON(400, "仅允许中文、英文字母、数字和空白字符，和常见标点符号")
		return
	}
	//地址中文校验
	if check.CheckChinese(&customer.Address) {
		c.JSON(400, "仅允许中文、英文字母、数字和空白字符，和常见标点符号")
		return
	}
	//手机号码校验
	if check.CheckPhoneNumber(&customer.Phone) {
		c.JSON(400, "手机号码输入非法")
		return

	}
	//todo 密码加密
	//所有校验通过
	err := dao.Add(&customer)
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
		u        Customer
		values   []Customer
		password string
	)
	u.Phone = c.Query("phone")
	password = c.Query("password")
	if u.Phone == "" || password == "" {
		c.JSON(400, "输入参数不能为空")
		return
	}
	//手机号有效性检验
	if check.CheckPhoneNumber(&u.Phone) {
		c.JSON(400, "手机号码校验非法")
		return

	}
	//对应账户查询
	err := dao.PerfectMatch(&u, &values)
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

// Delete 根据 id 删除指定用户，及其相关信息
// @Tags 用户管理
// @Summary 根据 id 删除指定用户，及其相关信息
// @Description 根据 id 删除指定用户，及其相关信息，包括其订单和收藏的菜品
// @Produce json
// @Param id query int true "用户id"
// @Success 200 {object} string "删除成功"
// @Success 400 {object} string "输入非法"
// @Success 500 {object} string "删除失败"
// @Router /customer [delete]
func (customer Customer) Delete(c *gin.Context) {

	id, err := strconv.Atoi(c.Query("id"))
	if id == 0 || err != nil {
		c.JSON(400, nil)
		return
	}

	customer = Customer{
		ID: uint(id),
	}
	//级联删除用户信息，包括其订单和收藏的菜品
	//主要是为了保护用户隐私
	err = dao.Del(&customer, 3)
	delCheck(c, err)

}
