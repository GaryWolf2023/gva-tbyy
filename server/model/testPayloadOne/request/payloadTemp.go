package request

import "os"

type CreateTemp struct {
	TempName        string  `json:"tempName"`
	TempVersion     string  `json:"tempVersion"`
	TempType        string  `json:"tempType"` // 辨别模板类型x-emr,form,markdown
	TempContentType string  `json:"tempContentType"`
	CreatePerson    string  `json:"createPerson"`
	UpdatePerson    string  `json:"updatePerson"`
	Active          bool    `json:"active"` // 判断是否被存档
	TempFile        os.File `json:"tempFile"`
}

type UpdateTemp struct {
	ID         string     `json:"id"`
	CreateTemp CreateTemp `json:"temp"`
}

type GetTempList struct {
	SearchData string
	Page       int
	PageSize   int
}

type DeleteTemp struct {
	ID string `json:"id"`
}
