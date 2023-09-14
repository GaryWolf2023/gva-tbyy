package request

type CreatePayload struct {
	Order          string `json:"order"`
	PayloadContent string `json:"payload_content" bind:"required"`
	JobId          string `json:"job_id" bind:"required"`
	Unique         string `json:"unique" bind:"required"`
	Business       string `json:"business" bind:"required"`
	DocName        string `json:"doc_name" bind:"required"`
	DocType        string `json:"doc_type" bind:"required"`
}

type UpdatePayload struct {
	PayloadId      string `json:"payload_id" bind:"required"`
	PayloadContent string `json:"payload_content" bind:"required"`
	SaveHis        string `json:"save_his" bind:"required"`
}

type DeletePayload struct {
	PayloadId string `json:"payload_id" bind:"required"`
}
type GetPayloadById struct {
	PayloadId string `json:"payload_id" bind:"required"`
}

type GetPayloadList struct {
	JobId  string `json:"job_id"`
	Unique string `json:"unique" bind:"required"`
}
