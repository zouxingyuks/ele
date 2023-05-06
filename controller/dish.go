package controller

import (
	"ele/models"
	"ele/tools"
	"ele/tools/dao"
	"github.com/gin-gonic/gin"
	"strconv"
)

//todo 提示信息大修

// AddDish 添加菜品
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
func AddDish(c *gin.Context) {
	var d models.Dish
	c.ShouldBind(&d)
	//中文校验

	err := tools.CheckChinese(&d.Name)
	if err != nil {
		c.JSON(400, "菜品名称"+err.Error())
		return

	}
	//todo 价格限定
	//todo 图片扫描
	//中文校验
	err = tools.CheckChinese(&d.Description)
	if err != nil {
		c.JSON(400, d.Description+err.Error())
		return

	}
	//店铺存在性检验
	var merchants []models.Merchant
	err = dao.PerfectMatch(&models.Merchant{ID: d.MerchantID}, &merchants)
	if err != nil {
		c.JSON(400, "餐厅不存在")
		return
	}
	//所有校验通过
	err = dao.Add(&d)
	addCheck(c, err)
}

// UpdateDish 更新菜品信息
// @Tags 商家管理
// @Summary 更新商家信息
// @Description 更新商家信息
// @Accept multipart/form-data
// @Produce multipart/json
// @Param id formData int true "菜品 id"
// @Param name formData string false "菜品名称"
// @Param description formData string false "菜品描述"
// @Param price formData int true "菜品价格"

// @Param phone formData string false "餐厅电话"
// @Success 200 {object} interface{} "添加成功"
// @Success 400 {object} string "添加失败"
// @Success 401 {object} string "输入非法"
// @Router /merchants [put]
func UpdateDish(c *gin.Context) {
	var (
		m   models.Merchant
		err error
	)
	_ = c.ShouldBind(&m)
	//中文校验
	if m.Name != "" {
		err = tools.CheckChinese(&m.Name)
		if err != nil {
			c.JSON(401, "店铺名称"+err.Error())
			return

		}
	}
	if m.Address != "" {
		//中文校验
		err = tools.CheckChinese(&m.Address)
		if err != nil {
			c.JSON(401, m.Address+err.Error())
			return

		}
	}
	//手机号码校验
	if m.Phone != "" {
		err = tools.CheckPhoneNumber(&m.Phone)
		if err != nil {
			c.JSON(401, m.Phone+err.Error())
			return

		}
	}
	mOld := models.Merchant{ID: m.ID}
	var values []models.Merchant
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

// ListDish 列出所有菜品
// @Tags 菜品管理
// @Summary 列出所有菜品
// @Description 获取所有菜品列表
// @Produce application/json
// @Success 200 {array} interface{} "Dish"
// @Failure 404 {object} string "资源未找到"
// @Failure 500 {object} string "查询失败"
// @Router /dishes [get]
func ListDish(c *gin.Context) {
	var values []models.Dish
	err := dao.List(&values)
	findCheck(c, values, err)
}

// PerfectDish 准确获取菜品信息
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
func PerfectDish(c *gin.Context) {
	d := models.Dish{
		Name: c.Query("name"),
	}
	merchantID, err := strconv.Atoi(c.Query("merchantID"))
	d.MerchantID = uint(merchantID)
	if d.Name == "" || d.MerchantID == 0 {
		c.JSON(400, "输入参数不能为空")
		return
	}

	var values []models.Dish
	err = dao.PerfectMatch(&d, &values, "Comments")
	findCheck(c, values, err)
}

// FuzzyDish 模糊搜索菜品信息
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
func FuzzyDish(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		c.JSON(400, "请求参数不能为空")
		return
	}
	var values []models.Dish
	err := dao.FuzzyMatch(name, &values, "Comments")
	findCheck(c, values, err)
}

// DeleteDish 根据 id 删除指定菜品,及其评论
// @Tags 菜品管理
// @Summary 根据 id 删除指定菜品,及其评论
// @Description 根据 id 删除指定菜品,及其评论
// @Produce json
// @Param id query int true "菜品id"
// @Success 200 {object} string "删除成功"
// @Success 400 {object} string "输入非法"
// @Success 500 {object} string "删除失败"
// @Router /dishes [delete]
func DeleteDish(c *gin.Context) {

	id, err := strconv.Atoi(c.Query("id"))
	if id == 0 || err != nil {
		c.JSON(400, "输入非法")
		return
	}

	d := models.Dish{ID: uint(id)}
	//永久删除菜品，选择级联硬删除
	err = dao.Del(&d, 3)
	delCheck(c, err)

}
