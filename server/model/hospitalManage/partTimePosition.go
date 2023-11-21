package hospitalManage

type HrEmployeePersonJob struct {
	EmployeeId int `json:"employeeId" form:"employeeId" gorm:"column:employee_id"`
	PositionId int `json:"positionId" form:"positionId" gorm:"column:position_id"`
}

func (h *HrEmployeePersonJob) Table() string {
	return "hr_employee_person_job"
}
