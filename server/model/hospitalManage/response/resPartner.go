package response

type AddressList struct {
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
}
