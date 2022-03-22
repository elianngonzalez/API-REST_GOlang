package commons

import (
	"API-REST/models"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func GetConection() *gorm.DB {
	dns := "root:@/test?parseTime=true"
	db, err := gorm.Open("mysql", dns)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func Migrate() {
	db := GetConection()
	defer db.Close()

	log.Println("Migrando...")

	db.AutoMigrate(&models.Persona{})
}
