package response

type PayloadDoc struct {
	Header       string `json:"header"`
	PayloadId    string `json:"payload_id"`
	Active       bool   `json:"active"`
	Meta         string `json:"meta"`
	Participants string `json:"participants"`
	Payload      string `json:"payload"`
}

type PayloadList struct {
	PayloadId string `json:"payload_id"`
}
