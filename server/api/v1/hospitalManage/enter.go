package hospitalManage

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	StaffManage
	WorkForceManage
}

var (
	StaffService = service.ServiceGroupApp.HospitalManageServiceGroup.StaffManageService
)
