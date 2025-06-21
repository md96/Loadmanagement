package db

import (
	//"load-management/models"
	"log"

	"github.com/md96/load-management/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBInit() {
	dsn := "root:root@tcp(localhost:3306)/recordings"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Issue while connecting to DB", err)
	}
	DB.AutoMigrate(&models.Gridstations{}, &models.Powerlog{})

}
