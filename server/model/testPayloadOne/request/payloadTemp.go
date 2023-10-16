package request

type CreateTemp struct {
	TempName        string `json:"tempName"`
	TempType        string `json:"tempType"` // 辨别模板类型x-emr,form,markdown
	TempContentType string `json:"tempContentType"`
	CreatePerson    string `json:"createPerson"`
	Active          bool   `json:"active"` // 判断是否被存档
	//TempFile        os.File `json:"tempFile"`
}

type UpdateTemp struct {
	TempID          int    `json:"id"`
	UpdatePerson    string `json:"updatePerson"` // 这个是人的id----后面改成int8
	TempName        string `json:"tempName"`
	TempType        string `json:"tempType"` // 辨别模板类型x-emr,form,markdown
	TempContentType string `json:"tempContentType"`
	Active          bool   `json:"active"` // 判断是否被存档
}

type UpdateOfFile struct {
	ID              string `json:"id"`
	BucketName      string `json:"bucketName"`
	ObjectName      string `json:"objectName"`
	OldContent      string `json:"oldContent"`
	PayloadTemplate string `json:"payloadTemplate"`
	Version         int8   `json:"version"`
}

type GetTempList struct {
	Type       string `json:"type"`
	SearchData string `json:"searchData"`
	Page       int    `json:"page"`
	PageSize   int    `json:"pageSize"`
}

type GetTempFile struct {
	ID       int    `form:"id"`
	TempName string `form:"tempName"`
	TempType string `form:"tempType" bind:"required"`
}

type GetFile struct {
	BucketName string `form:"bucketName"`
	ObjectName string `form:"objectName"`
}

type TempContent struct {
	Bucket string
	Key    string
}

type DeleteTemp struct {
	ID string `json:"id"`
}
