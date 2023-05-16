package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zouxingyuks/tools/check"
	"github.com/zouxingyuks/tools/dao"
	"gorm.io/gorm"
	"strconv"
	"time"
)

// Dish 菜品
type Dish struct {
	ID          uint `gorm:"primarykey" json:"id,omitempty" form:"id" binding:"required"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Name        string         `gorm:"type:varchar(255);comment:菜品名称" json:"name,omitempty" form:"name" binding:"required"`       // 菜品名称
	Description string         `gorm:"type:text;comment:菜品描述" json:"description,omitempty" form:"description" binding:"required"` // 菜品描述
	Price       float64        `gorm:"type:double;comment:菜品价格" json:"price,omitempty" form:"price" binding:"required"`           // 菜品价格
	//todo 图片改成数组，一个菜品可以有多张图片
	Picture    string    `gorm:"type:varchar(255);comment:菜品图片" json:"picture,omitempty" form:"picture" binding:"required"` // 菜品图片
	MerchantID uint      `gorm:"comment:所属餐厅id" json:"merchantID,omitempty" form:"merchantID" binding:"required"`           // 所属餐厅id
	Comments   []Comment // 一个菜品有很多评价，使用外键关联
}

//todo 提示信息大修

// Add 添加菜品
// @Tags 菜品管理
// @Summary 添加菜品
// @Description 添加菜品
// @Accept multipart/form-data
// @Produce application/json
// @Param name formData string true "菜品名称"
// @Param description formData string true "菜品描述"
// @Param picture formData string  false "菜品图片"
// @Param price formData float64 true "菜品价格"
// @Param merchantID formData uint true "所属餐厅id"
// @Success 200 {object} string "添加成功"
// @Success 400 {object} string "输入非法"
// @Success 500 {object} string "添加失败"
// @Router /dishes [post]
func (d Dish) Add(c *gin.Context) {
	c.ShouldBind(&d)
	//中文校验

	if check.CheckChinese(&d.Name) {
		c.JSON(400, "仅允许中文、英文字母、数字和空白字符，和常见标点符号")
		return

	}
	//todo 价格限定
	//todo 图片扫描
	//中文校验
	if check.CheckChinese(&d.Description) {
		c.JSON(400, "仅允许中文、英文字母、数字和空白字符，和常见标点符号")
		return

	}
	//店铺存在性检验
	var merchants []Merchant
	err := dao.PerfectMatch(&Merchant{ID: d.MerchantID}, &merchants)
	if err != nil {
		c.JSON(400, "餐厅不存在")
		return
	}
	//所有校验通过
	err = dao.Add(&d)
	addCheck(c, err)
}

// Update 更新菜品信息
// @Tags 菜品管理
// @Summary 更新菜品信息
// @Description 更新菜品信息
// @Accept multipart/form-data
// @Produce multipart/json
// @Param id formData int true "菜品 id"
// @Param name formData string false "菜品名称"
// @Param description formData string false "菜品描述"
// @Param price formData float64 false "菜品价格"
// @Param picture formData string false "菜品图片"
// @Param phone formData string false "餐厅电话"
// @Success 200 {object} interface{} "添加成功"
// @Success 400 {object} string "添加失败"
// @Success 401 {object} string "输入非法"
// @Router /dishes [put]
func (d Dish) Update(c *gin.Context) {
	var (
		err error
	)
	_ = c.ShouldBind(&d)
	//中文校验
	if d.Name != "" {
		if check.CheckChinese(&d.Name) {
			c.JSON(401, "仅允许中文、英文字母、数字和空白字符，和常见标点符号")
			return

		}
	}
	//中文校验
	if d.Description != "" {
		if check.CheckChinese(&d.Description) {
			c.JSON(401, "仅允许中文、英文字母、数字和空白字符，和常见标点符号")
			return

		}
	}

	dOld := Dish{ID: d.ID}
	var values []Dish
	err = dao.PerfectMatch(&dOld, &values, "Dishes")
	if err != nil && err.Error() != "" {
		c.JSON(500, err)

	} else {
		if len(values) == 0 {
			//资源未找到
			c.JSON(404, "对应餐厅不存在")
			return
		}
		dNew := values[0]
		if d.Name != "" {
			dNew.Name = d.Name
		}
		if d.Description != "" {
			dNew.Description = d.Description
		}
		if d.Price != 0 {
			dNew.Price = d.Price
		}
		if d.Picture != "" {
			dNew.Picture = d.Picture
		}

		err := dao.Update(&dNew)
		if err != nil {
			//todo ????
			return
		}
		c.JSON(200, nil)
	}

}

// List 列出所有菜品
// @Tags 菜品管理
// @Summary 列出所有菜品
// @Description 获取所有菜品列表
// @Produce application/json
// @Success 200 {array} interface{} "Dish"
// @Failure 404 {object} string "资源未找到"
// @Failure 500 {object} string "查询失败"
// @Router /dishes [get]
func (d Dish) List(c *gin.Context) {
	var values []Dish
	err := dao.List(&values)
	findCheck(c, values, err)
}

// Perfect 准确获取菜品信息
// @Tags 菜品管理
// @Summary 准确获取菜品信息
// @Description 根据菜品名称准确获取菜品信息
// @Produce application/json
// @Param name query string true "菜品名称"
// @Param merchantID query uint true "所属餐厅id"
// @Success 200 {array} interface{} "Dish"
// @Failure 400 {object} string "输入参数不能为空"
// @Failure 404 {object} string "请求资源不存在"
// @Failure 500 {object} string "查询失败"
// @Router /dishes/perfect [get]
func (d Dish) Perfect(c *gin.Context) {
	d = Dish{
		Name: c.Query("name"),
	}
	merchantID, err := strconv.Atoi(c.Query("merchantID"))
	d.MerchantID = uint(merchantID)
	if d.Name == "" || d.MerchantID == 0 {
		c.JSON(400, "输入参数不能为空")
		return
	}

	var values []Dish
	err = dao.PerfectMatch(&d, &values, "Comments")
	findCheck(c, values, err)
}

// Fuzzy 模糊搜索菜品信息
// @Tags 菜品管理
// @Summary 模糊搜索菜品信息
// @Description 根据菜品名称模糊搜索菜品信息
// @Produce application/json
// @Param name query string true "菜品名称"
// @Success 200 {array} interface{} "Dish"
// @Failure 400 {object} string "请求参数不能为空"
// @Failure 404 {object} string "请求资源不存在"
// @Failure 500 {object} string "查询失败"
// @Router /dishes/fuzzy [get]
func (d Dish) Fuzzy(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		c.JSON(400, "请求参数不能为空")
		return
	}
	var values []Dish
	//todo sql 过滤
	err := dao.FuzzyMatch(name, &values, "Comments")
	findCheck(c, values, err)
}

// Delete 根据 id 删除指定菜品,及其评论
// @Tags 菜品管理
// @Summary 根据 id 删除指定菜品,及其评论
// @Description 根据 id 删除指定菜品,及其评论
// @Produce json
// @Param id query int true "菜品id"
// @Success 200 {object} string "删除成功"
// @Success 400 {object} string "输入非法"
// @Success 500 {object} string "删除失败"
// @Router /dishes [delete]
func (d Dish) Delete(c *gin.Context) {

	id, err := strconv.Atoi(c.Query("id"))
	if id == 0 || err != nil {
		c.JSON(400, "输入非法")
		return
	}

	d = Dish{ID: uint(id)}
	//永久删除菜品，选择级联硬删除
	err = dao.Del(&d, 3)
	delCheck(c, err)

}
