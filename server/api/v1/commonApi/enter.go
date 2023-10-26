package commonApi

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	GetTableDataApi
}

var (
	CommonApiService = service.ServiceGroupApp.CommonApiServiceGroup.CommonApiService
)
