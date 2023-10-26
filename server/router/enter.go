package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/commonApi"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/hospitalManage"
	"github.com/flipped-aurora/gin-vue-admin/server/router/participant"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
	"github.com/flipped-aurora/gin-vue-admin/server/router/testPayload"
)

type RouterGroup struct {
	System         system.RouterGroup
	Example        example.RouterGroup
	TestPayload    testPayload.RouterGroup
	Participant    participant.RouterGroup
	HospitalManage hospitalManage.RouterGroup
	CommonApi      commonApi.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
