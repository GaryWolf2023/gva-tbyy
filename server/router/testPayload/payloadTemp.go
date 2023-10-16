package testPayload

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type PayloadTempRouter struct{}

func (p *PayloadTempRouter) InitPayloadTempRouter(Router *gin.RouterGroup, RouterPub *gin.RouterGroup) {
	payloadTempRouter := Router.Group("payload").Use(middleware.OperationRecord())
	payloadTempApi := v1.ApiGroupApp.TestPayloadApiGroup.PayloadTemplate
	{
		payloadTempRouter.POST("createTemp", payloadTempApi.CreatePayloadTemp)
		payloadTempRouter.PUT("updateTemp", payloadTempApi.UpdatePayloadTemp)
		payloadTempRouter.DELETE("deleteTemp", payloadTempApi.DeletePayloadTemp)
		payloadTempRouter.POST("getTempList", payloadTempApi.GetPayloadTempList)
		payloadTempRouter.GET("getTemp", payloadTempApi.GetPayloadTemp)
		payloadTempRouter.POST("updatePayloadOfFile", payloadTempApi.UpdatePayloadOfFile)
		payloadTempRouter.GET("getFileOfTemp", payloadTempApi.GetFileOfTemp)
		payloadTempRouter.GET("getFile", payloadTempApi.GetFile)
	}
}
