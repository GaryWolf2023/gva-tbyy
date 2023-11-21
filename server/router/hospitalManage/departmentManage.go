package hospitalManage

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type DepartmentManageRouter struct{}

func (d *DepartmentManageRouter) InitDepartmentManageRouter(Router *gin.RouterGroup, RouterPub *gin.RouterGroup) {
	departmentManageRouter := Router.Group("department").Use(middleware.OperationRecord())

	depMagApi := v1.ApiGroupApp.HospitalManageApiGroup.DepartmentManageApi
	{
		departmentManageRouter.POST("createDep", depMagApi.CreateDep)
		departmentManageRouter.POST("updateDep", depMagApi.UpdateDep)
		departmentManageRouter.DELETE("deleteDep", depMagApi.DeleteDep)
		departmentManageRouter.GET("getDepList", depMagApi.GetDepList)
		departmentManageRouter.POST("batchedDel", depMagApi.BatchesOfDelete)
	}
}
