package request

type GetIllnessList struct {
	EmployeeId int `json:"employeeId" form:"employeeId"`
}
type AddIllness struct {
	EmployeeId int `json:"employeeId"`
	IllnessId  int `json:"illnessId"`
}
type DeleteIllness struct {
	EmployeeId int `form:"employeeId"`
	IllnessId  int `form:"illnessId"`
}
