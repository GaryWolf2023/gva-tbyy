package hospitalManage

type RouterGroup struct {
	StaffManageRouter
	StaffAdeptIllnessManageRouter
	JobTitle
	PrescriptionQualification
	BasicInsuranceRouter
	PartTimePositionRouter
	DepartmentManageRouter
	CompanyManageRouter
	AddressManageRouter
}
