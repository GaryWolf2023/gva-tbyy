package hospitalManage

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	StaffManage
	StaffAdeptIllnessApi
	WorkForceManage
	StaffInsuranceApi
	QualificationManageApi
	JobTitleApi
	PrescriptionQualification
	ManagePartTimePositionApi
	DepartmentManageApi
	CompanyManageApi
	AddressManageApi
}

var (
	StaffService            = service.ServiceGroupApp.HospitalManageServiceGroup.StaffManageService
	AdeptIllnessService     = service.ServiceGroupApp.HospitalManageServiceGroup.AdeptIllnessService
	InsuranceService        = service.ServiceGroupApp.HospitalManageServiceGroup.InsuranceManageService
	JobTitleService         = service.ServiceGroupApp.HospitalManageServiceGroup.JobTitleService
	PPQService              = service.ServiceGroupApp.HospitalManageServiceGroup.PrescriptionQualificationService
	PartTimePositionService = service.ServiceGroupApp.HospitalManageServiceGroup.PartTimePositionService
	DepartmentManageService = service.ServiceGroupApp.HospitalManageServiceGroup.DepartmentManageService
	CompanyManageService    = service.ServiceGroupApp.HospitalManageServiceGroup.CompanyManageService
	AddressManageService    = service.ServiceGroupApp.HospitalManageServiceGroup.AddressManageService
)
