package hospitalManage

import (
	"time"
)

type HospitalPerformanceCalled struct {
	ID           int       `json:"id"`
	CreateUid    int       `json:"createUid"`
	WriteUid     int       `json:"writeUid"`
	CreateDate   time.Time `json:"createDate" gorm:"column:create_date"`
	WriteDate    time.Time `json:"writeDate" gorm:"column:write_date"`
	Orderno      int       `json:"orderno"`
	Name         string    `json:"name"`
	Ceilingvalue string    `json:"ceilingvalue"`
	Floorvalue   string    `json:"floorvalue"`
	Category     string    `json:"category"`
	Note         string    `json:"note"`
	IndicatValue float64   `json:"indicatValue"`
	Grade        string    `json:"grade"`
}
type CreateJT struct {
	CreateUid    int       `json:"createUid"`
	WriteUid     int       `json:"writeUid"`
	CreateDate   time.Time `json:"createDate" gorm:"column:create_date"`
	WriteDate    time.Time `json:"writeDate" gorm:"column:write_date"`
	Orderno      int       `json:"orderno"`
	Name         string    `json:"name"`
	Ceilingvalue string    `json:"ceilingvalue"`
	Floorvalue   string    `json:"floorvalue"`
	Category     string    `json:"category"`
	Note         string    `json:"note"`
	IndicatValue float64   `json:"indicatValue"`
	Grade        string    `json:"grade"`
}
