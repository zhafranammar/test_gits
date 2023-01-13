package seed

import (
	"log"

	"github.com/jinzhu/gorm"

	"github.com/zhafranammar/rest-api/api/models"
)


func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.User{}, &models.Author{},&models.Publisher{},&models.Book{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}, &models.Author{},&models.Publisher{},&models.Book{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}
}