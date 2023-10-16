package system

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type MinioRouter struct{}

func (s *MenuRouter) InitMinioRouter(Router *gin.RouterGroup, RouterPub *gin.RouterGroup) (R gin.IRoutes) {
	minioRouter := Router.Group("minio").Use(middleware.OperationRecord())
	//minioRouterWithoutRecord := Router.Group("menu")
	sysMinioApi := v1.ApiGroupApp.SystemApiGroup.SysTemMinioApi
	{
		minioRouter.POST("updateFile", sysMinioApi.UpdateFileToMinio) // 上传文件
		minioRouter.GET("getFile", sysMinioApi.GetFileToMinio)        // 获取文件
		minioRouter.DELETE("deleteFile", sysMinioApi.DeleteToMinio)   // 删除文件
	}
	{
		// minioRouterWithoutRecord.POST("getMenu", authorityMenuApi.GetMenu)         // 获取菜单树
		// minioRouterWithoutRecord.POST("getMenuList", authorityMenuApi.GetMenuList) // 分页获取基础menu列表
	}
	return minioRouter
}
