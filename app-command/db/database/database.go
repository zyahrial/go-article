package database

import (
	"github.com/jinzhu/gorm"
	"fmt"
	"log"
	// "gorm.io/driver/postgres"
 	// "gorm.io/gorm"
)

var (
	DBCon *gorm.DB
)

func InitDB() {
	var err error

	// DBCon, err = gorm.Open(postgres.New(postgres.Config{
	// 	DSN: "user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Asia/jakarta",
	// 	PreferSimpleProtocol: true, // disables implicit prepared statement usage
	// }), &gorm.Config{})

	connectionString := "host=%s port=%s user=%s password=%s dbname=%s sslmode=%s"
	connectionString = fmt.Sprintf(connectionString, "postgres", "5432", "postgres", "postgres", "postgres", "disable")
    DBCon, err = gorm.Open("postgres", connectionString)
    if err != nil {
        log.Fatal(err)
    }
}