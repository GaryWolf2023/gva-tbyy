package request

type CreateAddress struct {
	EmployeeId  int    `json:"employeeId"`
	DisplayName string `json:"displayName"` // 名称
	Phone       string `json:"phone"`       // 电话
	Mobile      string `json:"mobile"`      // 手机
	Email       string `json:"email"`       // 邮箱
	Website     string `json:"website"`     // 网站
	Lang        string `json:"lang"`        // 语言
	UserId      int    `json:"userId"`      // 销售人
	City        int    `json:"city"`        // 城市
	CountryId   int    `json:"countryId"`   // 国家ID
	CompanyId   int    `json:"companyId"`   // 公司ID
	StateId     int    `json:"stateId"`     // 状态
	Vat         int    `json:"vat"`         // 税ID
	CategoryId  int    `json:"categoryId"`  // 标签ID
	//ActivityIds
}
type UpdateAddress struct {
	EmployeeId  int    `json:"employeeId"`
	ID          int    `json:"id"`
	DisplayName string `json:"displayName"` // 名称
	Phone       string `json:"phone"`       // 电话
	Email       string `json:"email"`       // 邮箱
	UserId      int    `json:"userId"`      // 销售人
	City        int    `json:"city"`        // 城市
	CountryId   int    `json:"countryId"`   // 国家ID
	CompanyId   int    `json:"companyId"`   // 公司ID
	StateId     int    `json:"stateId"`     // 状态
	CategoryId  int    `json:"categoryId"`  // 标签ID
	Vat         int    `json:"vat"`         // 税ID
	//ActivityIds
}
type DeleteAddress struct {
	ID int `form:"id"`
}
