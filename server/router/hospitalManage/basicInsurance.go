package hospitalManage

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type BasicInsuranceRouter struct{}

func (s *StaffManageRouter) InitBasicInsuranceRouter(Router *gin.RouterGroup, RouterPub *gin.RouterGroup) {
	basicInsuranceRouter := Router.Group("insurance").Use(middleware.OperationRecord())

	insuranceApi := v1.ApiGroupApp.HospitalManageApiGroup.StaffInsuranceApi
	{
		basicInsuranceRouter.POST("createInsurance", insuranceApi.CreateInsurance)
		basicInsuranceRouter.POST("updateInsurance", insuranceApi.UpdateInsurance)
		basicInsuranceRouter.DELETE("deleteInsurance", insuranceApi.DeleteInsurance)
		basicInsuranceRouter.GET("getInsuranceInfo", insuranceApi.GetInsuranceInfo)
		basicInsuranceRouter.GET("getInsuranceList", insuranceApi.GetInsuranceList)
		basicInsuranceRouter.GET("getPersonInsuranceList", insuranceApi.GetPersonInsuranceList)
		basicInsuranceRouter.POST("addPersonInsurance", insuranceApi.AddPersonInsurance)
		basicInsuranceRouter.POST("deletePersonInsurance", insuranceApi.DeletePersonInsurance)
	}
}
