package hospitalManage

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type StaffAdeptIllnessManageRouter struct{}

func (s *StaffManageRouter) InitStaffAdeptIllnessManageRouter(Router *gin.RouterGroup, RouterPub *gin.RouterGroup) {
	staffAdeptIllnessRouter := Router.Group("adeptIllness").Use(middleware.OperationRecord())

	staffAdeptIllnessApi := v1.ApiGroupApp.HospitalManageApiGroup.StaffAdeptIllnessApi
	{
		staffAdeptIllnessRouter.GET("getAdeptIllnessList", staffAdeptIllnessApi.GetAdeptIllnessList)
		staffAdeptIllnessRouter.POST("addAdeptIllness", staffAdeptIllnessApi.AddAdeptIllnessItem)
		staffAdeptIllnessRouter.DELETE("deleteAdeptIllness", staffAdeptIllnessApi.DeleteAdeptIllnessItem)
	}
}
