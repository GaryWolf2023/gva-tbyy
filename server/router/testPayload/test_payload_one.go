package testPayload

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TestPayloadOneRouter struct{}

func (t *TestPayloadOneRouter) InitoneRouter(Router *gin.RouterGroup, RouterPub *gin.RouterGroup) {
	oneRouter := Router.Group("payload").Use(middleware.OperationRecord())
	// oneRouterWithoutRecord := Router.Group("payload")

	//onePublicRouterWithoutRecord := RouterPub.Group("payload")
	oneRouterApi := v1.ApiGroupApp.TestPayloadApiGroup.TestPayloadOne
	{
		oneRouter.POST("createPayload", oneRouterApi.CreatePayload) // 创建Payload
		oneRouter.POST("updatePayload", oneRouterApi.UpdatePayload)
		oneRouter.POST("deletePayload", oneRouterApi.DeletePayload)
		oneRouter.POST("getPayloadById", oneRouterApi.GetPayloadById)
		oneRouter.POST("getPayloadList", oneRouterApi.GetPayloadList)
	}
	{
		// oneRouterWithoutRecord.POST("getAllApis", oneRouterApi.GetAllApis) // 获取所有api
		// oneRouterWithoutRecord.POST("getApiList", oneRouterApi.GetApiList) // 获取Api列表
	}
}
