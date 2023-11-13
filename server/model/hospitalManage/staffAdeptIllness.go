package hospitalManage

type HrEmployeeAdeptIllness struct {
	EmployeeId int `json:"employeeId" gorm:"column:employee_id"`
	IllnessId  int `json:"illnessId" gorm:"column:illness_id"`
}

func (h *HrEmployeeAdeptIllness) Table() string {
	return "hr_employee_adept_illness"
}
