package participant

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ParticipantRouter struct{}

func (p *ParticipantRouter) InitParticipantRouter(Router *gin.RouterGroup, RouterPub *gin.RouterGroup) {
	participantRouter := Router.Group("participant").Use(middleware.OperationRecord())
	participantApi := v1.ApiGroupApp.ParticipantApiGroup.ParticipantApi
	{
		participantRouter.POST("/participant/createParticipant", participantApi.CreateParticipant)
	}
}
