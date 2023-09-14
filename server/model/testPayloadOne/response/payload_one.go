package response

type PayloadDoc struct {
	PayloadId    string `json:"payload_id"`
	Active       bool   `json:"active"`
	Header       string `json:"header"`
	Meta         string `json:"meta"`
	Participants string `json:"participants"`
	Payload      string `json:"payload"`
}
