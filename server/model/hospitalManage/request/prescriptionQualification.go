package request

type GetList struct {
}
type CreatePQ struct {
	EmployeeId int    `json:"employeeId"`
	Name       string `json:"name"`
	Active     bool   `json:"active"`
}
type UpdatePQ struct {
	ID         uint   `json:"id"`
	EmployeeId int    `json:"employeeId"`
	Name       string `json:"name"`
	Active     bool   `json:"active"`
}
type DeletePQ struct {
	ID int `form:"id"`
}
type GetListPPQ struct {
	EmployeeId int `form:"employeeId"`
}
type AddPPQ struct {
	EmployeeId     int `json:"employeeId"`
	PrescriptionId int `json:"prescriptionId"`
}
type DeletePPQ struct {
	EmployeeId     int `form:"employeeId"`
	PrescriptionId int `form:"prescriptionId"`
}
