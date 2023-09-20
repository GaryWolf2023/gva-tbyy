package participant

type CreateParticipant struct {
	Name            string `json:"name"`
	ParticipantType string `json:"participant_type"`
	Remark          string `json:"remark"`
}
