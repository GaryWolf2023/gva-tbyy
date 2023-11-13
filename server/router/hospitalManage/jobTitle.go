package hospitalManage

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type JobTitle struct{}

// 职称管理
func (j *JobTitle) InitJobTitleRouter(Router *gin.RouterGroup, RouterPub *gin.RouterGroup) {
	jobTitleRouter := Router.Group("jobTitle").Use(middleware.OperationRecord())
	JobTitleApi := v1.ApiGroupApp.HospitalManageApiGroup.JobTitleApi
	{
		jobTitleRouter.GET("getList", JobTitleApi.GetJobTitleList)
		jobTitleRouter.POST("create", JobTitleApi.CreateJobTitle)
		jobTitleRouter.POST("update", JobTitleApi.UpdateJobTitle)
		jobTitleRouter.DELETE("delete", JobTitleApi.DeleteJobTitle)
	}
}
