package response

import (
	"github.com/shopspring/decimal"
	"time"
)

type GetInsuranceList struct {
	ID            int64           `json:"id"`
	CreateDate    time.Time       `json:"createDate"`
	UpdateDate    time.Time       `json:"updateDate"`
	CreateUid     int64           `json:"createUid"`
	WriteUid      int64           `json:"writeUid"`
	Name          string          `json:"name"`          // 保险类型名称
	BaseNum       decimal.Decimal `json:"baseNum"`       // 缴费基数
	PercentUnit   decimal.Decimal `json:"percentUnit"`   // 单位缴纳百分比
	PercentPerson decimal.Decimal `json:"percentPerson"` // 个人缴纳百分比
}

type GetPersonInsuranceList struct {
	ID            int64           `json:"id"`
	EmployeeId    int64           `json:"employeeId"`
	Name          string          `json:"name"`          // 保险类型名称
	BaseNum       decimal.Decimal `json:"baseNum"`       // 缴费基数
	PercentUnit   decimal.Decimal `json:"percentUnit"`   // 单位缴纳百分比
	PercentPerson decimal.Decimal `json:"percentPerson"` // 个人缴纳百分比
}

type GetPersonInsurance struct {
	InsuranceId int64 `json:"insuranceId"`
	EmployeeId  int64 `json:"employeeId"`
}
