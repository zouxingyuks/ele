package controller

import (
	"ele/models"
	"ele/tools"
	"ele/tools/dao"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// AddDish 添加菜品
// @Tags 菜品管理
// @Summary 添加菜品
// @Description 添加菜品
// @Accept multipart/form-data
// @Produce json
// @Param name formData string true "菜品名称"
// @Param description formData string true "菜品描述"
// @Param picture formData string  false "菜品图片"
// @Param price formData double true "菜品价格"
// @Param merchantID formData int32 true "所属餐厅id"
// @Success 200 {object} models.Response "添加成功"
// @Success 400 {object} models.Response "添加失败"
// @Success 401 {object} models.Response "输入非法"
// @Router /dish/add [post]
func AddDish(c *gin.Context) {
	var d models.Dish
	fmt.Println(c.ContentType())

	c.ShouldBind(&d)
	//中文校验
	fmt.Println(2)

	err := tools.CheckChinese(&d.Name)
	if err != nil {
		c.JSON(401, models.Response{
			Msg:  "输入非法",
			Data: "菜品名称" + err.Error(),
		})
		return

	}
	fmt.Println(3)

	//todo 价格限定
	//todo 图片扫描
	//中文校验
	err = tools.CheckChinese(&d.Description)
	if err != nil {
		c.JSON(401, models.Response{
			Msg:  "输入非法",
			Data: d.Description + err.Error(),
		})
		return

	}
	fmt.Println(4)

	value := models.Merchant{}
	err = dao.PerfectMatch(&models.Merchant{
		Model: gorm.Model{
			ID: d.MerchantID,
		},
	}, &value)
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
		c.JSON(400, models.Response{
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
// @Produce json
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
// @Param name query string true "菜品名称"
// @Success 200 {object} models.Response "Dish"
// @Failure 400 {object} models.Response "ErrorResponse"
// @Failure 500 {object} models.Response "ErrorResponse"
// @Router /dish/perfect [get]
func PerfectDish(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		c.JSON(400, models.Response{
			Msg:  "请求参数不能为空",
			Data: nil,
		})
		return
	}

	value := models.Dish{}
	err := dao.PerfectMatch(&models.Dish{Name: name}, &value, "Comments")
	if err != nil {
		c.JSON(500, models.Response{
			Msg:  "获取失败",
			Data: err,
		})

	} else {
		c.JSON(200, models.Response{
			Msg:  "获取成功",
			Data: value,
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
// @Router /dish/fuzzy [get]
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
	c.JSON(200, models.Response{
		Msg: "获取成功",
		//Data: string(jsonMarshalData),
		Data: values,
	})

}
