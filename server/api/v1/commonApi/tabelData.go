package commonApi

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
)

type GetTableDataApi struct{}

// GetNation 获取民族
func (g *GetTableDataApi) GetNation(c *gin.Context) {
	var res request.GetTableNation
	err := c.ShouldBindQuery(&res)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
}

// GetCountry 获取省份
func (g *GetTableDataApi) GetCountry(c *gin.Context) {
	data := CommonApiService.GetTableCountry()
	response.OkWithDetailed(data, "获取成功", c)
}

// GetProvince 获取省份
func (g *GetTableDataApi) GetProvince(c *gin.Context) {
	var req request.GetTableProvince
	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	data := CommonApiService.GetTableProvince(req)
	response.OkWithDetailed(data, "获取成功", c)
}

// GetCity 获取市区
func (g *GetTableDataApi) GetCity(c *gin.Context) {
	var req request.GetTableCity
	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	data := CommonApiService.GetTableCity(req)
	response.OkWithDetailed(data, "获取成功", c)
}

// GetArea 获取区县
func (g *GetTableDataApi) GetArea(c *gin.Context) {
	var req request.GetTableArea
	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	data := CommonApiService.GetTableArea(req)
	response.OkWithDetailed(data, "获取成功", c)
}

// GetPolitics 获取政治面貌
func (g *GetTableDataApi) GetPolitics(c *gin.Context) {

}

// GetPolitics 获取政治面貌
func (g *GetTableDataApi) GetNumberCodeData(c *gin.Context) {
	var req request.GetNumberCode
	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	data := CommonApiService.GetNumberCode(req)
	response.OkWithDetailed(data, "获取成功", c)
}
