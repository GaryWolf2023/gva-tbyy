package participant

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TempManagePatRouter struct{}

func (t *TempManagePatRouter) InitTempManageRouter(Router *gin.RouterGroup, RouterPub *gin.RouterGroup) {
	tempManageRouter := Router.Group("participant").Use(middleware.OperationRecord())

	tempManageApi := v1.ApiGroupApp.ParticipantApiGroup.TempManageParticipant
	{
		tempManageRouter.POST("createTemp", tempManageApi.CreateTempOfParticipant)
		tempManageRouter.POST("updateTemp", tempManageApi.UpdateTempOfParticipant)
		tempManageRouter.GET("getTempList", tempManageApi.GetTempOfParticipant)
		tempManageRouter.DELETE("deleteTemp", tempManageApi.DeleteTempOfParticipant)
	}
}
