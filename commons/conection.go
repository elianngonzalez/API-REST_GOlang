package commons

import (
	"log"

	"github.com/jinzhu/gorm"
)

func GetConection() *gorm.DB {
	db, err := gorm.Open("mysql", "root:root@test?charset=utf8")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
