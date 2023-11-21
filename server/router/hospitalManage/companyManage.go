package hospitalManage

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CompanyManageRouter struct{}

func (c *CompanyManageRouter) InitCompanyManageRouter(Router *gin.RouterGroup, RouterPub *gin.RouterGroup) {
	companyManageRouter := Router.Group("company").Use(middleware.OperationRecord())
	companyManageApi := v1.ApiGroupApp.HospitalManageApiGroup.CompanyManageApi
	{
		companyManageRouter.GET("getList", companyManageApi.GetCompany)
		companyManageRouter.POST("create", companyManageApi.CreateCompany)
		companyManageRouter.POST("update", companyManageApi.UpdateCompany)
		companyManageRouter.DELETE("delete", companyManageApi.DeleteCompany)
	}
}
