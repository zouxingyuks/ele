package controller

import "github.com/gin-gonic/gin"

// @Summary 列出所有商家
// @Description 获取所有商家列表
// @Produce json
// @Success 200 {array} Merchant
// @Failure 500 {object} ErrorResponse
// @Router /merchants [get]
func ListMerchants(c *gin.Context) {

}

// @Summary 获取商家信息
// @Description 根据 ID 获取商家信息
// @Produce json
// @Param merchant_id path int true "商家 ID"
// @Success 200 {object} Merchant
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /merchants/{merchant_id} [get]
func GetMerchant(c *gin.Context) {

}
