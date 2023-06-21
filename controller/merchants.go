package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zouxingyuks/tools/check"
	"github.com/zouxingyuks/tools/dao"
	"gorm.io/gorm"
	"strconv"
	"time"
)

// Merchant 餐厅
type Merchant struct {
	ID        uint `gorm:"primarykey" json:"id,omitempty" form:"id" binding:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string         `gorm:"type:varchar(255);comment:餐厅名称" json:"name,omitempty" form:"name" binding:"required"`        // 餐厅名称
	Address   string         `gorm:"type:varchar(255);comment:餐厅地址" json:"address,omitempty" form:"address" binding:"required" ` // 餐厅地址
	Phone     string         `gorm:"type:varchar(20);comment:注册电话" json:"phone,omitempty" form:"phone" binding:"required"`       // 注册电话
	//todo 添加联系方式的结构体
	Dishes []Dish // 一个餐厅有多个菜品，使用外键关联
}

// Add 添加商家
// @Tags 商家管理
// @Summary 添加商家
// @Description 添加商家
// @Accept multipart/form-data
// @Produce multipart/json
// @Param name formData string true "餐厅名称"
// @Param address formData string true "餐厅地址"
// @Param phone formData string true "餐厅电话"
// @Success 200 {object} interface{} "添加成功"
// @Success 400 {object} string "添加失败"
// @Success 401 {object} string "输入非法"
// @Router /merchants [post]
func (m Merchant) Add(c *gin.Context) {
	_ = c.ShouldBind(&m)
	//中文校验
	if check.CheckChinese(&m.Name) {
		c.JSON(401, "仅允许中文、英文字母、数字和空白字符，和常见标点符号")
		return

	}
	//中文校验
	if check.CheckChinese(&m.Address) {
		c.JSON(401, "仅允许中文、英文字母、数字和空白字符，和常见标点符号")
		return

	}
	//手机号码校验
	if check.CheckPhoneNumber(&m.Phone) {
		c.JSON(401, "手机号码输入非法")
		return

	}

	err := dao.Add(&m)
	if err != nil {
		c.JSON(400, err)
		return
	}
	c.JSON(200, m)
	return
}

// Update 更新商家信息
// @Tags 商家管理
// @Summary 更新商家信息
// @Description 更新商家信息
// @Accept multipart/form-data
// @Produce multipart/json
// @Param id formData int true "餐厅 id"
// @Param name formData string false "餐厅名称"
// @Param address formData string false "餐厅地址"
// @Param phone formData string false "餐厅电话"
// @Success 200 {object} interface{} "添加成功"
// @Success 400 {object} string "添加失败"
// @Success 401 {object} string "输入非法"
// @Router /merchants [put]
func (m Merchant) Update(c *gin.Context) {
	var err error
	_ = c.ShouldBind(&m)
	//中文校验
	if m.Name != "" {
		if check.CheckChinese(&m.Name) {
			c.JSON(401, "仅允许中文、英文字母、数字和空白字符，和常见标点符号")
			return

		}
	}
	if m.Address != "" {
		//中文校验
		if check.CheckChinese(&m.Address) {
			c.JSON(401, "仅允许中文、英文字母、数字和空白字符，和常见标点符号")
			return

		}
	}
	//手机号码校验
	if m.Phone != "" {
		if check.CheckPhoneNumber(&m.Phone) {
			c.JSON(401, "手机号码输入非法")
			return

		}
	}
	mOld := Merchant{ID: m.ID}
	var values []Merchant
	err = dao.PerfectMatch(&mOld, &values, "Dishes")
	if err != nil && err.Error() != "" {
		c.JSON(500, err)

	} else {
		if len(values) == 0 {
			//资源未找到
			c.JSON(404, "对应餐厅不存在")
			return
		}
		mNew := values[0]
		if m.Name != "" {
			mNew.Name = m.Name
		}
		if m.Phone != "" {
			mNew.Phone = m.Phone
		}
		if m.Address != "" {
			mNew.Address = m.Address
		}
		err := dao.Update(&mNew)
		if err != nil {
			//todo ????
			return
		}
		c.JSON(200, nil)
	}
}

// List 列出所有商家
// @Tags 商家管理
// @Summary 列出所有商家
// @Description 获取所有商家列表
// @Produce multipart/json
// @Success 200 {array} interface{} "获取成功"
// @Failure 404 {object} string "资源未找到"
// @Failure 500 {object} string "查询失败"
// @Router /merchants [get]
func (m Merchant) List(c *gin.Context) {
	var values []Merchant
	err := dao.List(&values, "Dishes")
	findCheck(c, values, err)
}

// Perfect 准确获取商家信息
// @Tags 商家管理
// @Summary 准确获取商家信息
// @Description 根据商家名称获取商家信息
// @Produce multipart/json
// @Param name query string true "商家名称"
// @Success 200 {array} interface{} "获取成功"
// @Failure 400 {object} string "输入非法"
// @Failure 404 {object} string "资源未找到"
// @Failure 500 {object} string "查询失败"
// @Router /merchants/perfect [get]
func (m Merchant) Perfect(c *gin.Context) {
	m = Merchant{
		Name: c.Query("name"),
	}
	if m.Name == "" {
		c.JSON(400, "输入非法")
		return
	}

	var values []Merchant
	err := dao.PerfectMatch(&m, &values, "Dishes")
	findCheck(c, values, err)

}

// Fuzzy 模糊搜索商家信息
// @Tags 商家管理
// @Summary 模糊搜索商家信息
// @Description 根据商家名称模糊搜索商家信息
// @Produce multipart/json
// @Param name query string true "商家名称"
// @Success 200 {array} interface{} "获取成功"
// @Failure 400 {object} string "输入非法"
// @Failure 404 {object} string "资源未找到"
// @Failure 500 {object} string "查询失败"
// @Router /merchants/fuzzy [get]
func (m Merchant) Fuzzy(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		c.JSON(400, nil)
		return
	}
	var values []Merchant
	err := dao.FuzzyMatch(name, &values, "Dishes")
	findCheck(c, values, err)
}

// Delete 根据 id 删除指定商家，及其菜品
// @Tags 商家管理
// @Summary 根据 id 删除指定商家，及其菜品
// @Description 根据 id 删除指定商家，及其菜品
// @Produce multipart/json
// @Param id query int true "商家id"
// @Success 200 {object} string "删除成功"
// @Success 400 {object} string "输入非法"
// @Success 500 {object} string "删除失败"
// @Router /merchants [delete]
func (m Merchant) Delete(c *gin.Context) {

	id, err := strconv.Atoi(c.Query("id"))
	if id == 0 || err != nil {
		c.JSON(400, nil)
		return
	}

	merchant := Merchant{ID: uint(id)}
	//永久删除商家，选择级联硬删除
	err = dao.Del(&merchant, 3)
	//todo 伪bug 已经 删除的能够再次删除
	delCheck(c, err)

}
