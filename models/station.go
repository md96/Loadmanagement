package models

type Gridstations struct {
	ID       int        `json:"id" gorm:"primaryKey"`
	Status   string     `json:"status"`
	Name     string     `json:"name"`
	Powerlog []Powerlog `json:"-" gorm:"foreignKey:StationId"`
}

type Powerlog struct {
	ID        uint `gorm:"primaryKey"`
	StationId int  `gorm:"index"`
	Power     float64
}
