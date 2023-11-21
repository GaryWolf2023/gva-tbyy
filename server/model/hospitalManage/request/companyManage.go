package request

/*
*创建公司，只是需要创建基础信息
 */
type CreateCompany struct {
	EmployeeId int `json:"employeeId" gorm:"column:employee_id"`
}

/*
*关于更新这部分需要更多的兼容性
*可能只是单个更新，个可能同时更新多列，太复杂了,能够使用form直接更新
 */
type UpdateCompany struct {
	ID int `json:"id"`
}

type DeleteCompany struct {
	ID int `json:"id"`
}
