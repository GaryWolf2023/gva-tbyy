package response

import "time"

type GetTemp struct {
	ID              int       `json:"id"`
	CreateDate      time.Time `json:"createDate"`
	UpdateDate      time.Time `json:"updateDate"`
	TempName        string    `json:"tempName"`
	TempVersion     int8      `json:"tempVersion"`
	TempContent     string    `json:"tempContent"`
	TempType        string    `json:"tempType"` // 辨别模板类型x-emr,form,markdown
	TempContentType string    `json:"tempContentType"`
	CreatePerson    string    `json:"createPerson"`
	UpdatePerson    string    `json:"updatePerson"`
	Active          bool      `json:"active"` // 判断是否被存档
}
