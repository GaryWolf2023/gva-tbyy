package response

type GetTemp struct {
	ID              string `json:"id"`
	TempName        string `json:"tempName"`
	TempVersion     string `json:"tempVersion"`
	TempContent     string `json:"tempContent"`
	TempType        string `json:"tempType"` // 辨别模板类型x-emr,form,markdown
	TempContentType string `json:"tempContentType"`
	CreatePerson    string `json:"createPerson"`
	UpdatePerson    string `json:"updatePerson"`
	Active          bool   `json:"active"` // 判断是否被存档
}
