package commonApi

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type TableDataRouter struct{}

func (t *TableDataRouter) InitTableDataRouter(PubRouter *gin.RouterGroup) {
	commonApiRouter := PubRouter.Group("common")
	commonApi := v1.ApiGroupApp.CommonApiApiGroup.GetTableDataApi
	{
		commonApiRouter.GET("nation", commonApi.GetNation)
		commonApiRouter.GET("country", commonApi.GetCountry)
		commonApiRouter.GET("province", commonApi.GetProvince)
		commonApiRouter.GET("city", commonApi.GetCity)
		commonApiRouter.GET("area", commonApi.GetArea)
		commonApiRouter.GET("numberCode", commonApi.GetNumberCodeData)
	}
}
