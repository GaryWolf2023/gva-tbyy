package response

type Info struct {
	Id         int    `json:"id"`
	ParentId   int    `json:"parentId"` // 父部门ID
	Name       string `json:"name"`
	EmployeeId int    `json:"employeeId"`
	Note       string `json:"note"`      // 描述
	Property   int    `json:"property"`  // 科室属性
	DutyId     int    `json:"dutyId"`    // 负责人ID
	Code       string `json:"code"`      // 科室编码-不能重复
	Virtual    bool   `json:"virtual"`   // 是否实体部门
	Available  bool   `json:"available"` // 是否可用
}
