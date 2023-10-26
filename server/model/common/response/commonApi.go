package response

type Country struct {
	ID   int    `json:"value"`
	Name string `json:"label"`
}
type Province struct {
	ID   int    `json:"value"`
	Name string `json:"label"`
}
type City struct {
	ID   int    `json:"value"`
	Name string `json:"label"`
}
type Area struct {
	ID   int    `json:"value"`
	Name string `json:"label"`
}
type NumCodeData struct {
	ValueMean string `json:"label"`
	ID        int64  `json:"value"`
}
