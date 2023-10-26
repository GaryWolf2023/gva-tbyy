package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/commonApi"
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/hospitalManage"
	"github.com/flipped-aurora/gin-vue-admin/server/service/participant"
	"github.com/flipped-aurora/gin-vue-admin/server/service/payloadOne"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup         system.ServiceGroup
	ExampleServiceGroup        example.ServiceGroup
	PayloadServiceGroup        payloadOne.ServiceGroup
	ParticipantServiceGroup    participant.ServiceGroup
	HospitalManageServiceGroup hospitalManage.ServiceGroup
	CommonApiServiceGroup      commonApi.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
