package testpayloadone

import "github.com/flipped-aurora/gin-vue-admin/server/global"

type PayloadTemplate struct {
	global.GVA_MODEL
	TempName    string
	TempVersion string
	Temp
}
