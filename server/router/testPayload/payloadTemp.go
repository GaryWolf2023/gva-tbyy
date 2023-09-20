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
		payloadTempRouter.POST("createTemp", payloadTempApi.GetPayloadTemp)
		payloadTempRouter.PUT("updateTemp", payloadTempApi.UpdatePayloadTemp)
		payloadTempRouter.DELETE("deleteTemp", payloadTempApi.DeletePayloadTemp)
		payloadTempRouter.GET("getTemp", payloadTempApi.GetPayloadTemp)
	}
}
