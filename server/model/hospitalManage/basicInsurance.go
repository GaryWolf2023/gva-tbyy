package hospitalManage

import (
	"github.com/shopspring/decimal"
	"time"
)

type HrEmployeeInsurance struct {
	ID            int64           `json:"id" gorm:"column:id"`
	CreateDate    time.Time       `json:"createDate" gorm:"column:create_date"`
	UpdateDate    time.Time       `json:"updateDate" gorm:"column:update_date"`
	CreateUid     int64           `json:"createUid" gorm:"column:create_uid"`
	WriteUid      int64           `json:"writeUid" gorm:"column:write_uid"`
	Name          string          `json:"name" gorm:"column:name"`                    // 保险类型名称
	BaseNum       decimal.Decimal `json:"baseNum" gorm:"column:base_num"`             // 缴费基数
	PercentUnit   decimal.Decimal `json:"percentUnit" gorm:"column:percent_unit"`     // 单位缴纳百分比
	PercentPerson decimal.Decimal `json:"percentPerson" gorm:"column:percent_person"` // 个人缴纳百分比
}

type HrEmployeePersonInsurances struct {
	InsuranceId int64 `json:"insuranceId" gorm:"insurance_id"`
	EmployeeId  int64 `json:"employeeId" gorm:"employee_id"`
}

func (h *HrEmployeeInsurance) table() string {
	return "hr_employee_insurance"
}
func (e *HrEmployeePersonInsurances) table() string {
	return "hr_employee_person_insurances"
}
