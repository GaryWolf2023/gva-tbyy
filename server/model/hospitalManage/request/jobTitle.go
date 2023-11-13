package request

type GetJobTitle struct {
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
type CreateJobTitle struct {
	EmployeeId   int     `json:"employeeId"`
	Name         string  `json:"name"`
	Ceilingvalue string  `json:"ceilingvalue"`
	Floorvalue   string  `json:"floorvalue"`
	Category     string  `json:"category"`
	Note         string  `json:"note"`
	IndicatValue float64 `json:"indicatValue"`
	Grade        string  `json:"grade"`
}

type UpdateJobTitle struct {
	ID           int     `json:"id"`
	EmployeeId   int     `json:"employeeId"`
	Name         string  `json:"name"`
	Ceilingvalue string  `json:"ceilingvalue"`
	Floorvalue   string  `json:"floorvalue"`
	Category     string  `json:"category"`
	Note         string  `json:"note"`
	IndicatValue float64 `json:"indicatValue"`
	Grade        string  `json:"grade"`
}
type DeleteJobTitle struct {
	ID int `json:"id"`
}
