package controller

import (
	"ele/models"
	"ele/tools"
	"ele/tools/dao"
	"github.com/gin-gonic/gin"
	"strconv"
)

// AddMerchant 添加商家
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
func AddMerchant(c *gin.Context) {
	var m models.Merchant
	_ = c.ShouldBind(&m)
	//中文校验
	err := tools.CheckChinese(&m.Name)
	if err != nil {
		c.JSON(401, "店铺名称"+err.Error())
		return

	}
	//中文校验
	err = tools.CheckChinese(&m.Address)
	if err != nil {
		c.JSON(401, m.Address+err.Error())
		return

	}
	//手机号码校验
	err = tools.CheckPhoneNumber(&m.Phone)
	if err != nil {
		c.JSON(401, m.Phone+err.Error())
		return

	}

	err = dao.Add(&m)
	if err != nil {
		c.JSON(400, err)
		return
	}
	c.JSON(200, m)
	return

}

// UpdateMerchant 更新商家信息
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
func UpdateMerchant(c *gin.Context) {
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

// ListMerchant 列出所有商家
// @Tags 商家管理
// @Summary 列出所有商家
// @Description 获取所有商家列表
// @Produce multipart/json
// @Success 200 {array} interface{} "获取成功"
// @Failure 404 {object} string "资源未找到"
// @Failure 500 {object} string "查询失败"
// @Router /merchants [get]
func ListMerchant(c *gin.Context) {
	var values []models.Merchant
	err := dao.List(&values, "Dishes")
	findCheck(c, values, err)
}

// PerfectMerchant 准确获取商家信息
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
func PerfectMerchant(c *gin.Context) {
	m := models.Merchant{
		Name: c.Query("name"),
	}
	if m.Name == "" {
		c.JSON(400, "输入非法")
		return
	}

	var values []models.Merchant
	err := dao.PerfectMatch(&m, &values, "Dishes")
	findCheck(c, values, err)

}

// FuzzyMerchant 模糊搜索商家信息
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
func FuzzyMerchant(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		c.JSON(400, nil)
		return
	}
	var values []models.Merchant
	err := dao.FuzzyMatch(name, &values, "Dishes")
	findCheck(c, values, err)
}

// DeleteMerchant 根据 id 删除指定商家，及其菜品
// @Tags 商家管理
// @Summary 根据 id 删除指定商家，及其菜品
// @Description 根据 id 删除指定商家，及其菜品
// @Produce multipart/json
// @Param id query int true "商家id"
// @Success 200 {object} string "删除成功"
// @Success 400 {object} string "输入非法"
// @Success 500 {object} string "删除失败"
// @Router /merchants [delete]
func DeleteMerchant(c *gin.Context) {

	id, err := strconv.Atoi(c.Query("id"))
	if id == 0 || err != nil {
		c.JSON(400, nil)
		return
	}

	merchant := models.Merchant{ID: uint(id)}
	//永久删除商家，选择级联硬删除
	err = dao.Del(&merchant, 3)
	//todo 伪bug 已经 删除的能够再次删除
	delCheck(c, err)

}
