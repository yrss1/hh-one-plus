package database

import (
	"2-ass/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

var DB *gorm.DB

func Init() *gorm.DB {
	dsn := "host=localhost user=postgres password=postgres dbname=postgres sslmode=disable "
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(models.Coin{})
	return db
}
func GetDB() *gorm.DB {
	if DB == nil {
		DB = Init()
		var sleep = time.Duration(1)
		for DB == nil {
			sleep = sleep * 2
			fmt.Printf("Database in unavalible. Wait for %d sec.\n", sleep)
			time.Sleep(sleep * time.Second)
			DB = Init()
		}
	}
	return DB
}
