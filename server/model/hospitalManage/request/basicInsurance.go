package request

import (
	"github.com/shopspring/decimal"
	"time"
)

type CreateInsurance struct {
	CreateDate    time.Time       `json:"createDate"`
	CreateUid     int64           `json:"createUid"`
	Name          string          `json:"name"`          // 保险类型名称
	BaseNum       decimal.Decimal `json:"baseNum"`       // 缴费基数
	PercentUnit   decimal.Decimal `json:"percentUnit"`   // 单位缴纳百分比
	PercentPerson decimal.Decimal `json:"percentPerson"` // 个人缴纳百分比
}
type UpdateInsurance struct {
	ID            int64           `json:"id"`
	UpdateDate    time.Time       `json:"updateDate"`
	WriteUid      int64           `json:"writeUid"`
	Name          string          `json:"name"`          // 保险类型名称
	BaseNum       decimal.Decimal `json:"baseNum"`       // 缴费基数
	PercentUnit   decimal.Decimal `json:"percentUnit"`   // 单位缴纳百分比
	PercentPerson decimal.Decimal `json:"percentPerson"` // 个人缴纳百分比
}
type DeleteInsurance struct {
	ID int64 `json:"id" form:"id"`
}

type GetInsuranceInfo struct {
	ID int64 `json:"id" form:"id"`
}

type GetPersonInsuranceList struct {
	ID int64 `json:"id" form:"id"` // employeeId
}

type GetPersonInsurance struct {
	InsuranceId int64 `json:"insuranceId"`
	EmployeeId  int64 `json:"employeeId"`
}
type AddPersonInsurance struct {
	InsuranceId int64 `json:"insuranceId"`
	EmployeeId  int64 `json:"employeeId"`
}

type DeletePersonInsurance struct {
	InsuranceId int64 `json:"insuranceId"`
	EmployeeId  int64 `json:"employeeId"`
}
