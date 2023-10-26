package commonApi

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
)

type CommonApiService struct{}

func (c *CommonApiService) GetTableNation(req request.GetTableNation) interface{} {
	//global.GVA_DB.Table().
	return nil
}

func (c *CommonApiService) GetTableCountry() interface{} {
	var country []response.Country
	global.GVA_DB.Table("res_country").Find(&country)
	return country
}

func (c *CommonApiService) GetTableProvince(req request.GetTableProvince) interface{} {
	var province []response.Province
	global.GVA_DB.Table("res_country_state").Where("country_id = ?", req.CountryID).Select([]string{"id", "name"}).Find(&province)
	fmt.Println(province)
	return province
}

func (c *CommonApiService) GetTableCity(req request.GetTableCity) interface{} {
	var city []response.City
	fmt.Println(req.ProvinceID)
	global.GVA_DB.Table("res_country_city").Where("state_id = ?", req.ProvinceID).Select([]string{"id", "name"}).Find(&city)
	return city
}
func (c *CommonApiService) GetTableArea(req request.GetTableArea) interface{} {
	var area []response.Area
	global.GVA_DB.Table("res_country_city").Where("city_id = ?", req.CityID).Select([]string{"id", "name"}).Find(&area)
	return area
}
func (c *CommonApiService) GetNumberCode(req request.GetNumberCode) interface{} {
	var numberCode []response.NumCodeData
	global.GVA_DB.Table("numbercode").Where("code = ?", req.Code).Find(&numberCode)
	return numberCode
}
