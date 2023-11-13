package hospitalManage

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type StaffManageRouter struct{}

func (s *StaffManageRouter) InitStaffManageRouter(Router *gin.RouterGroup, RouterPub *gin.RouterGroup) {
	staffManageRouter := Router.Group("staff").Use(middleware.OperationRecord())

	staffManageApi := v1.ApiGroupApp.HospitalManageApiGroup.StaffManage
	{
		staffManageRouter.GET("getStaffList", staffManageApi.GetStaffInfoList)
		staffManageRouter.GET("getStaff", staffManageApi.GetStaffInfo)
		staffManageRouter.POST("createStaff", staffManageApi.CreateStaff)
		staffManageRouter.POST("batchCreate", staffManageApi.GetStaffInfoList)
		staffManageRouter.DELETE("deleteStaff", staffManageApi.DeleteStaff)
		staffManageRouter.DELETE("batchDelete", staffManageApi.GetStaffInfoList)
		staffManageRouter.POST("updateStaff", staffManageApi.UpdateStaffInfo)
		staffManageRouter.POST("updateBasicStaff", staffManageApi.UpdateBasicStaffInfo)
		staffManageRouter.POST("updateSystemConfig", staffManageApi.UpdateSystemConfig)
	}
}
