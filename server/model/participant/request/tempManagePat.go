package request

type CreateTempOfPat struct {
	ParticipantName    string `json:"participant_name" bind:"required"`
	ParticipantContent string `json:"participant_content" bind:"required"`
	BusinessType       string `json:"business_type" bind:"required"`
}

type UpdateTempOfPat struct {
	ParticipantID      int    `json:"participantId" bind:"required"`
	ParticipantName    string `json:"participantName" bind:"required"`
	ParticipantContent string `json:"participantContent" bind:"required"`
	BusinessType       string `json:"businessType" bind:"required"`
}
type DeleteTempOfPat struct {
	ParticipantID int `json:"participant_id" bind:"required"`
}

type GetTempOfPat struct {
	SearchData   string `json:"searchData"`
	BusinessType string `json:"businessType" bind:"required"`
}
