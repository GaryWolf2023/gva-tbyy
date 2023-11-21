package request

type GetItemList struct {
	EmployeeId int `form:"employeeId"`
}
type AddItem struct {
	EmployeeId int `json:"employeeId"`
	PositionId int `json:"positionId"`
}
type DeleteItem struct {
	EmployeeId int `form:"employeeId"`
	PositionId int `form:"positionId"`
}
type CreateJob struct {
	EmployeeId        int    `json:"employeeId"`
	Name              string `json:"name"`
	NoOfEmployee      int    `json:"noOfEmployee"`
	NoOfRecruitment   int    `json:"noOfRecruitment"`
	NoOfHiredEmployee int    `json:"noOfHiredEmployee"`
	UserId            int    `json:"userId"`
	Active            bool   `json:"active"`
	NaturePost        string `json:"naturePost"`
	CategoryPost      string `json:"categoryPost"`
	JobCode           string `json:"jobCode"`
	Show              string `json:"show"`
}
type BasicJobStruct struct {
	MessageMainAttachmentId int    `json:"messageMainAttachmentId" gorm:"column:message_main_attachment_id"`
	Sequence                int    `json:"sequence" gorm:"column:sequence"`
	ExpectedEmployees       int    `json:"expectedEmployees" gorm:"column:expected_employees"`
	NoOfEmployee            int    `json:"noOfEmployee" gorm:"column:no_of_employee"`
	NoOfRecruitment         int    `json:"noOfRecruitment" gorm:"column:no_of_recruitment"`
	NoOfHiredEmployee       int    `json:"noOfHiredEmployee" gorm:"column:no_of_hired_employee"`
	DepartmentId            int    `json:"departmentId" gorm:"column:department_id"`
	CompanyId               int    `json:"companyId" gorm:"column:company_id"`
	ContractTypeId          int    `json:"contractTypeId" gorm:"column:contract_type_id"`
	Name                    string `json:"name" gorm:"column:name"`
	Description             string `json:"description" gorm:"column:description"`
	Requirements            string `json:"requirements" gorm:"column:requirements"`
	Active                  bool   `json:"active" gorm:"column:active"`
	AliasId                 int    `json:"aliasId" gorm:"column:alias_id"`
	AddressId               int    `json:"addressId" gorm:"column:address_id"`
	ManagerId               int    `json:"managerId" gorm:"column:manager_id"`
	UserId                  int    `json:"userId" gorm:"column:user_id"`
	HrResponsibleId         int    `json:"hrResponsibleId" gorm:"column:hr_responsible_id"`
	Color                   int    `json:"color" gorm:"column:color"`
	NaturePost              string `json:"naturePost" gorm:"column:nature_post"`
	CategoryPost            string `json:"categoryPost" gorm:"column:category_post"`
	JobCode                 string `json:"jobCode" gorm:"column:job_code"`
	Show                    string `json:"show" gorm:"column:show"`
}
type UpdateJob struct {
	UserId       int            `json:"userId"`
	UpdateObject BasicJobStruct `json:"updateObject"`
}
type DeleteJob struct {
	ID int `json:"id"`
}
