package controller

import (
	"ele/models"
	"ele/tools"
	"ele/tools/dao"
	"fmt"
	"github.com/gin-gonic/gin"
)

// AddMerchant 添加商家
// @Tags 商家管理
// @Summary 添加商家
// @Description 添加商家
// @Accept multipart/form-data
// @Produce json
// @Param name formData string true "餐厅名称"
// @Param address formData string true "餐厅地址"
// @Param phone formData string true "餐厅电话"
// @Success 200 {object} models.Response "添加成功"
// @Success 400 {object} models.Response "添加失败"
// @Success 401 {object} models.Response "输入非法"
// @Router /merchant/add [post]
func AddMerchant(c *gin.Context) {
	var m models.Merchant
	c.ShouldBind(&m)
	//中文校验
	err := tools.CheckChinese(&m.Name)
	if err != nil {
		c.JSON(401, models.Response{
			Msg:  "输入非法",
			Data: "店铺名称" + err.Error(),
		})
		return

	}
	//中文校验
	err = tools.CheckChinese(&m.Address)
	if err != nil {
		c.JSON(401, models.Response{
			Msg:  "输入非法",
			Data: m.Address + err.Error(),
		})
		return

	}
	//手机号码校验
	err = tools.CheckPhoneNumber(&m.Phone)
	if err != nil {
		c.JSON(401, models.Response{
			Msg:  "输入非法",
			Data: m.Phone + err.Error(),
		})
		return

	}

	err = dao.Add(&m)
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

// ListMerchant 列出所有商家
// @Tags 商家管理
// @Summary 列出所有商家
// @Description 获取所有商家列表
// @Produce json
// @Success 200 {object} models.Response
// @Failure 500 {object} models.Response "ErrorResponse"
// @Router /merchant/list [get]
func ListMerchant(c *gin.Context) {
	var values []models.Merchant
	dao.List(&values, "Dishes")
	c.JSON(200, models.Response{
		Msg:  "下面是所有商家数据",
		Data: values,
	})
}

// PerfectMerchant 准确获取商家信息
// @Tags 商家管理
// @Summary 准确获取商家信息
// @Description 根据商家名称获取商家信息
// @Produce json
// @Param name query string true "商家名称"
// @Success 200 {object} models.Response "Merchant"
// @Failure 400 {object} models.Response "ErrorResponse"
// @Failure 500 {object} models.Response "ErrorResponse"
// @Router /merchant/perfect [get]
func PerfectMerchant(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		c.JSON(400, models.Response{
			Msg:  "请求参数不能为空",
			Data: nil,
		})
		return
	}

	value := models.Merchant{}
	err := dao.PerfectMatch(&models.Merchant{Name: name}, &value, "Dishes")
	//jsonMarshalData, _ := json.Marshal(value)
	fmt.Println(err)
	if err != nil {
		c.JSON(500, models.Response{
			Msg: "获取失败",
			//Data: string(jsonMarshalData),
			Data: err,
		})

	} else {
		c.JSON(200, models.Response{
			Msg: "获取成功",
			//Data: string(jsonMarshalData),
			Data: value,
		})
	}

}

// FuzzyMerchant 模糊搜索商家信息
// @Tags 商家管理
// @Summary 模糊搜索商家信息
// @Description 根据商家名称模糊搜索商家信息
// @Produce json
// @Param name query string true "商家名称"
// @Success 200 {object} models.Response "Merchant"
// @Failure 400 {object} models.Response "ErrorResponse"
// @Failure 500 {object} models.Response "ErrorResponse"
// @Router /merchant/fuzzy [get]
func FuzzyMerchant(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		c.JSON(400, models.Response{
			Msg:  "请求参数不能为空",
			Data: nil,
		})
		return
	}
	var values []models.Merchant
	dao.FuzzyMatch(name, &values, "Dishes")
	//jsonMarshalData, _ := json.Marshal(value)

	c.JSON(200, models.Response{
		Msg: "获取成功",
		//Data: string(jsonMarshalData),
		Data: values,
	})
}
