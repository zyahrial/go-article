package migrations

import (
	"command/article/db/database"
	"command/article/models"
	"log"
)

func Migrate() {

	err := database.DBCon.AutoMigrate(&models.Article{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

}