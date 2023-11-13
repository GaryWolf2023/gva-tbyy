package hospitalManage

import "github.com/flipped-aurora/gin-vue-admin/server/global"

type HrEmployeePersonPrescription struct {
	EmployeeId     int `json:"employeeId" gorm:"column:employee_id"`
	PrescriptionId int `json:"prescriptionId" gorm:"column:prescription_id"`
}
type HrEmployeeChufang struct {
	global.GVA_MODEL
	CreateUid int    `json:"createUid" gorm:"column:create_uid"`
	WriteUid  int    `json:"writeUid" gorm:"column:update_uid"`
	Name      string `json:"name" gorm:"column:name"`
	Active    bool   `json:"active" gorm:"column:active"`
}

func (h *HrEmployeePersonPrescription) Table() string {
	return "hr_employee_person_prescription"
}
func (h *HrEmployeeChufang) Table() string {
	return "hr_employee_chufang"
}
