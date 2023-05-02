package controller

import (
	"ele/models"
	"ele/tools"
	"ele/tools/dao"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

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
// @Success 200 {object} models.Response "添加成功"
// @Success 400 {object} models.Response "输入非法"
// @Success 500 {object} models.Response "添加失败"
// @Router /dish/add [post]
func AddDish(c *gin.Context) {
	var d models.Dish
	c.ShouldBind(&d)
	//中文校验

	err := tools.CheckChinese(&d.Name)
	if err != nil {
		c.JSON(400, models.Response{
			Msg:  "输入非法",
			Data: "菜品名称" + err.Error(),
		})
		return

	}
	//todo 价格限定
	//todo 图片扫描
	//中文校验
	err = tools.CheckChinese(&d.Description)
	if err != nil {
		c.JSON(400, models.Response{
			Msg:  "输入非法",
			Data: d.Description + err.Error(),
		})
		return

	}
	//店铺存在性检验
	var merchants []models.Merchant
	err = dao.PerfectMatch(&models.Merchant{
		Model: gorm.Model{
			ID: d.MerchantID,
		},
	}, &merchants)
	if err != nil {
		c.JSON(400, models.Response{
			Msg:  "添加失败",
			Data: "餐厅不存在",
		})
		return
	}
	//所有校验通过
	err = dao.Add(&d)
	if err != nil {
		c.JSON(500, models.Response{
			Msg:  "添加失败",
			Data: err,
		})
		return
	}
	c.JSON(200, models.Response{
		Msg:  "添加成功",
		Data: nil,
	})
	return

}

// ListDish 列出所有菜品
// @Tags 菜品管理
// @Summary 列出所有菜品
// @Description 获取所有菜品列表
// @Produce application/json
// @Success 200 {object} models.Response "获取成功"
// @Failure 500 {object} models.Response "ErrorResponse"
// @Router /dish/list [get]
func ListDish(c *gin.Context) {
	var values []models.Dish
	dao.List(&values)
	c.JSON(200, models.Response{
		Msg:  "下面是所有菜品信息",
		Data: values,
	})
}

// PerfectDish 准确获取菜品信息
// @Tags 菜品管理
// @Summary 准确获取菜品信息
// @Description 根据菜品名称准确获取菜品信息
// @Produce json
// @Param name formData string true "菜品名称"
// @Param merchantID formData uint false "所属餐厅id"
// @Success 200 {object} models.Response "Dish"
// @Failure 400 {object} models.Response "ErrorResponse"
// @Success 404 {object} models.Response "请求资源不存在"
// @Failure 500 {object} models.Response "ErrorResponse"
// @Router /dish/perfect [post]
func PerfectDish(c *gin.Context) {
	d := models.Dish{}
	c.ShouldBind(&d)
	if d.Name == "" || d.MerchantID == 0 {
		c.JSON(400, models.Response{
			Msg:  "请求参数不能为空",
			Data: nil,
		})
		return
	}

	var values []models.Dish
	err := dao.PerfectMatch(&d, &values, "Comments")
	if err != nil {
		c.JSON(500, models.Response{
			Msg:  "获取失败",
			Data: err,
		})

	} else {
		if len(values) == 0 {
			c.JSON(404, models.Response{
				Msg: "请求资源不存在",
				//Data: string(jsonMarshalData),
				Data: values,
			})
			return
		}
		c.JSON(200, models.Response{
			Msg:  "获取成功",
			Data: values,
		})
	}
}

// FuzzyDish 模糊搜索菜品信息
// @Tags 菜品管理
// @Summary 模糊搜索菜品信息
// @Description 根据菜品名称模糊搜索菜品信息
// @Produce json
// @Param name query string true "菜品名称"
// @Success 200 {object} models.Response "Dish"
// @Success 400 {object} models.Response "请求参数不能为空"
// @Success 404 {object} models.Response "请求资源不存在"
// @Router /dish/fuzzy [post]
func FuzzyDish(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		c.JSON(400, models.Response{
			Msg:  "请求参数不能为空",
			Data: nil,
		})
		return
	}
	var values []models.Dish
	dao.FuzzyMatch(name, &values, "Comments")
	//jsonMarshalData, _ := json.Marshal(value)
	if len(values) == 0 {
		c.JSON(404, models.Response{
			Msg: "请求资源不存在",
			//Data: string(jsonMarshalData),
			Data: values,
		})
		return
	}
	c.JSON(200, models.Response{
		Msg: "获取成功",
		//Data: string(jsonMarshalData),
		Data: values,
	})

}
