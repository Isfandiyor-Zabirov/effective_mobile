package database

import (
	"effective_mobile_tech_task/utils/env"
	"fmt"
	"github.com/kr/pretty"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var db *gorm.DB

func InitDB() {
	var err error
	settings := env.GetSettings()

	db, err = gorm.Open(postgres.Open(fmt.Sprintf("host=%s user=%s port=%s dbname=%s password=%s sslmode=disable",
		settings.Database.Host,
		settings.Database.Username,
		settings.Database.Port,
		settings.Database.Name,
		settings.Database.Password)),
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		},
	)

	if err != nil {
		log.Fatalf("Connect to database error: %s", err.Error())
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(100)
	pretty.Logln("DB successfully connected! ")
}

func GetDB() *gorm.DB {
	return db
}

// CloseDB closes database connection (unnecessary)
func CloseDB() {
	sqlDB, err := db.DB()
	sqlDB.Close()
	if err != nil {
		pretty.Logln("Error on closing the DB: ", err)
	}
}
