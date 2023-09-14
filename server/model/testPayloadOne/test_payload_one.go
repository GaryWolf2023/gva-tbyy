package testpayloadone

import "github.com/flipped-aurora/gin-vue-admin/server/global"

type PayloadDoc struct {
	global.GVA_MODEL
	PayloadId    string `json:"payload_id" gorm:"column:payload_id;primaryKey"`
	Active       bool   `json:"active" gorm:"column:active;"`
	Header       string `json:"header" gorm:"column:header"`
	Meta         string `json:"meta" gorm:"column:meta"`
	Participants string `json:"participants" gorm:"column:participants"`
	Payload      string `json:"payload" gorm:"column:payload"`
}

type Header struct {
	JobId     string
	Unique    string
	Business  string
	PayloadId string
}
type Meta struct {
	Version string
	DocName string
	DocType string
	NeedPay bool
	IsPay   bool
}
type PayloadContent struct {
	Order   string
	Content string
}
type Participent struct {
	Name string
}

func (p *PayloadDoc) TableName() string {
	return "payload_doc"
}
