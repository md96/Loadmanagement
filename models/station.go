package models

type Gridstations struct {
	ID       int        `json:"id" gorm:"primaryKey" example:1002 `
	Status   string     `json:"status" example:"Active"`
	Name     string     `json:"name"  example:"Main Grid Station"`
	Powerlog []Powerlog `json:"-" gorm:"foreignKey:StationId"`
}

type Powerlog struct {
	ID        uint `gorm:"primaryKey"`
	StationId int  `gorm:"index"`
	Power     float64
}
