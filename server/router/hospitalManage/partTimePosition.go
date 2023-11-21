package hospitalManage

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type PartTimePositionRouter struct{}

func (p *PartTimePositionRouter) InitPartTimePositionRouter(Router *gin.RouterGroup, RouterPub *gin.RouterGroup) {
	PartTiemPositionRouter := Router.Group("partTimePosition").Use(middleware.OperationRecord())
	PartTimePositionApi := v1.ApiGroupApp.HospitalManageApiGroup.ManagePartTimePositionApi
	{
		PartTiemPositionRouter.GET("getList", PartTimePositionApi.GetPartTimePositionList)
		PartTiemPositionRouter.POST("add", PartTimePositionApi.AddPartTimePosition)
		PartTiemPositionRouter.DELETE("delete", PartTimePositionApi.DeletePartTimePosition)

		PartTiemPositionRouter.GET("getJobList", PartTimePositionApi.GetJobList)
		PartTiemPositionRouter.POST("createJob", PartTimePositionApi.CreateJob)
		PartTiemPositionRouter.POST("updateJob", PartTimePositionApi.UpdateJob)
		PartTiemPositionRouter.DELETE("deleteJob", PartTimePositionApi.DeleteJob)
	}
}
