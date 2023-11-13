package request

type GetList struct {
	ID        int    `json:"id"`
	CreateUid int    `json:"createUid"`
	WriteUid  int    `json:"writeUid"`
	Name      string `json:"name"`
	Active    bool   `json:"active"`
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
	ID int `json:"id"`
}
type GetListPPQ struct {
	EmployeeId int `json:"employeeId"`
}
type AddPPQ struct {
	EmployeeId     int `json:"employeeId"`
	PrescriptionId int `json:"prescriptionId"`
}
type DeletePPQ struct {
	EmployeeId     int `json:"employeeId"`
	PrescriptionId int `json:"prescriptionId"`
}
