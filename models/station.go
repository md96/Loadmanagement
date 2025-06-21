package models

type Station struct {
	ID       int `gorm:"primaryKey"`
	Status   string
	Name     string
	Powerlog []Powerlog `gorm:"foreignKey:StationId"`
}

type Powerlog struct {
	ID        uint `gorm:"primaryKey"`
	StationId int  `gorm:"index"`
	Power     float64
}
