package testPayload

import "github.com/flipped-aurora/gin-vue-admin/server/service/payloadOne"

type ApiGroup struct {
	TestPayloadOne
	PayloadTemplate
}

var (
	PayloadService = new(payloadOne.PayloadService)
)
