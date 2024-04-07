package migration

import (
	"effective_mobile_tech_task/internal/entitites/cars"
	"effective_mobile_tech_task/logger"
	"effective_mobile_tech_task/pkg/database"
	"gorm.io/gorm"
	logger2 "gorm.io/gorm/logger"
)

func AutoMigrate() {
	db := database.GetDB()

	err := db.Session(&gorm.Session{Logger: logger2.Default.LogMode(logger2.Silent)}).
		AutoMigrate(&cars.People{}, &cars.Car{})
	if err != nil {
		logger.Error.Fatalf("error in migration: %s", err.Error())
	}
}
