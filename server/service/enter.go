package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/participant"
	"github.com/flipped-aurora/gin-vue-admin/server/service/payloadOne"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup      system.ServiceGroup
	ExampleServiceGroup     example.ServiceGroup
	PayloadServiceGroup     payloadOne.ServiceGroup
	ParticipantServiceGroup participant.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
