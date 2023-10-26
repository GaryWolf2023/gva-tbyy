package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/commonApi"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/hospitalManage"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/participant"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/testPayload"
)

type ApiGroup struct {
	SystemApiGroup         system.ApiGroup
	ExampleApiGroup        example.ApiGroup
	TestPayloadApiGroup    testPayload.ApiGroup
	ParticipantApiGroup    participant.ApiGroup
	HospitalManageApiGroup hospitalManage.ApiGroup
	CommonApiApiGroup      commonApi.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
