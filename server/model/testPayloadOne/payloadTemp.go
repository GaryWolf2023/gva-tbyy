package testpayloadone

import "github.com/flipped-aurora/gin-vue-admin/server/global"

// PayloadTemplate 模板文档
type PayloadTemplate struct {
	global.GVA_MODEL
	TempName        string `json:"tempName" gorm:"column:temp_name"`
	TempVersion     int8   `json:"tempVersion" gorm:"column:temp_version"`
	TempContent     string `json:"tempContent" gorm:"column:temp_content"`
	TempType        string `json:"tempType" gorm:"column:temp_type"` // 辨别模板类型x-emr,form,markdown
	TempContentType string `json:"tempContentType" gorm:"column:temp_content_type"`
	CreatePerson    string `json:"createPerson" gorm:"column:create_person"`
	UpdatePerson    string `json:"updatePerson" gorm:"column:update_person"`
	Active          bool   `json:"active" gorm:"column:active"` // 判断是否被存档
}

// HistoryTemplate 历史文档，历史文档将会被存档
type HistoryTemplate struct {
	global.GVA_MODEL
	TempID          int8   `json:"tempID" gorm:"column:temp_id;primaryKey;"`
	TempVersion     int8   `json:"tempVersion" gorm:"column:temp_version"`
	TempContent     string `json:"tempContent" gorm:"column:temp_content"`
	PayloadTemplate string `json:"payloadTemplate" gorm:"column:payload_template"`
}

type MinioSaveInfo struct {
	Bucket string
	Key    string
}

func (p *PayloadTemplate) Table() string {
	return "payload_templates"
}
func (h *HistoryTemplate) Table() string {
	return "payload_template_history"
}
