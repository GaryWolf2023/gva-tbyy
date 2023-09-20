package request

type GetParticipantsReq struct {
	Search string `json:"search"`
	Type   string `json:"type"`
}
