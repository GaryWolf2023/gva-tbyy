package response

type JobTitleList struct {
	ID           int     `json:"id"`
	Orderno      int     `json:"orderno"`
	Name         string  `json:"name"`
	Ceilingvalue string  `json:"ceilingvalue"`
	Floorvalue   string  `json:"floorvalue"`
	Category     string  `json:"category"`
	Note         string  `json:"note"`
	IndicatValue float64 `json:"indicatValue"`
	Grade        string  `json:"grade"`
}
