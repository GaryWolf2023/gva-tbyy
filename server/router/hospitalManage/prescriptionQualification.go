package hospitalManage

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type PrescriptionQualification struct{}

// 职称管理
func (j *JobTitle) InitPPRouter(Router *gin.RouterGroup, RouterPub *gin.RouterGroup) {
	ppRouter := Router.Group("prescription").Use(middleware.OperationRecord())
	ppApi := v1.ApiGroupApp.HospitalManageApiGroup.PrescriptionQualification
	{
		ppRouter.GET("getList", ppApi.GetListPQ)
		ppRouter.POST("create", ppApi.CreatePQ)
		ppRouter.POST("update", ppApi.UpdatePQ)
		ppRouter.DELETE("delete", ppApi.DeletePQ)
		ppRouter.POST("addPersonPQItem", ppApi.AddPersonPQItem)
		ppRouter.GET("getPersonPQList", ppApi.GetPersonPQList)
		ppRouter.DELETE("deletePersonPQ", ppApi.DeletePersonPQ)
	}
}
