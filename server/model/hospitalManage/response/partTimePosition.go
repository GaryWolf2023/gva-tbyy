package response

type PartTimePosition struct {
	EmployeeId int `json:"employeeId"`
	PositionId int `json:"positionId"`
}

// 获取个人兼职岗位工作列表
type PositionList struct {
	ID                int    `json:"id"`
	Sequence          int    `json:"sequence"`
	Name              string `json:"name"`
	JobCode           string `json:"jobCode"`           // 岗位代码
	NaturePost        string `json:"naturePost"`        // 岗位性质
	CategoryPost      string `json:"categoryPost"`      // 岗位类别
	Show              string `json:"show"`              // 岗位表述
	DepartmentId      int    `json:"departmentId"`      // 部门ID
	NoOfRecruitment   int    `json:"noOfRecruitment"`   // 目标员工数量
	NoOfEmployee      int    `json:"noOfEmployee"`      // 工作岗位现有员工数量
	UserId            int    `json:"userId"`            // 招聘人员
	Active            bool   `json:"active"`            // 是否启用
	AliasId           int    `json:"aliasId"`           // 别名ID
	ExpectedEmployees int    `json:"expectedEmployees"` // 此岗位期望员工数量
	NoOfHiredEmployee int    `json:"noOfHiredEmployee"` // 已招聘员工数量
	CompanyId         int    `json:"companyId"`         // 公司
}

// 尝试一次下直接返回interface{}
// 获取关于工作岗位列表
type GetJobList struct {
}
