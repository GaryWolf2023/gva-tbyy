package response

type GetList struct {
	ID        int    `json:"id"`
	CreateUid int    `json:"createUid"`
	WriteUid  int    `json:"writeUid"`
	Name      string `json:"name"`
	Active    bool   `json:"active"`
}
type GetPPQList struct {
	EmployeeId     int `json:"employeeId"`
	PrescriptionId int `json:"prescriptionId"`
}
