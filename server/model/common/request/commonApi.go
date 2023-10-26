package request

type GetTableProvince struct {
	CountryID string `json:"countryId" form:"countryId"`
}
type GetTableCity struct {
	ProvinceID string `json:"provinceId" form:"provinceId"`
}
type GetTableArea struct {
	CityID string `json:"cityId" form:"cityId"`
}
type GetTableNation struct {
	CountryID string `json:"countryId" form:"countryId"`
}
type GetNumberCode struct {
	Code string `json:"code" form:"code"`
}
