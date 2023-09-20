package participant

import "github.com/flipped-aurora/gin-vue-admin/server/global"

type ParticipantTemp struct {
	global.GVA_MODEL
	Activate           bool   `json:"activate" gorm:"activate"`
	ParticipantName    string `json:"participant_name" gorm:"participant_name"`
	ParticipantContent string `json:"participant_content" gorm:"participant_content"`
	BusinessType       string `json:"business_type" gorm:"business_type"`
}
type ParticipantItem struct {
	global.GVA_MODEL
	PatID      string
	PersonName string
	PersonType string
}

type ParticipantPersonType struct {
	//	参与人类型
	global.GVA_MODEL
	TypeID string `json:"type_id" gorm:"column:type_id;description:传入参与人类型的ID;"`
}

func (p *ParticipantTemp) Table() string {
	return "participant_temp"
}
func (p *ParticipantPersonType) Table() string {
	return "participant_person_type"
}
