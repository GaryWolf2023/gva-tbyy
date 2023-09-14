package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
	"github.com/flipped-aurora/gin-vue-admin/server/router/testPayload"
)

type RouterGroup struct {
	System      system.RouterGroup
	Example     example.RouterGroup
	TestPayload testPayload.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
