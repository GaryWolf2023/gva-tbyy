package response

type EmployeeAdeptIllness struct {
	EmployeeId int `json:"employeeId" gorm:"column:employee_id"`
	IllnessId  int `json:"illnessId" gorm:"column:illness_id"`
}
