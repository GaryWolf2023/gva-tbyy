package participant

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service"
)

type ApiGroup struct {
	ParticipantApi
	GetParticipantApi
	TempManageParticipant
}

var (
	PatTempService    = service.ServiceGroupApp.ParticipantServiceGroup.TempManagePatService
	PaticipantService = service.ServiceGroupApp.ParticipantServiceGroup.ParticipantService
	GetPatService     = service.ServiceGroupApp.ParticipantServiceGroup.GetParticipantService
)
