package main

import (
	"effective_mobile_tech_task/internal/api"
	"effective_mobile_tech_task/logger"
	"effective_mobile_tech_task/pkg/database"
	"effective_mobile_tech_task/pkg/migration"
	"effective_mobile_tech_task/utils/env"
)

// @title API документация
// @version 1.0.0
// @description Документация к сервису Tech project.
// @termsOfService http://swagger.io/terms/
// @contact.name Isfandiyor
// @contact.email isfand.zabirov@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @accept json
// @produce json
// @schemes http
func main() {
	env.ReadEnv()
	logger.Init()
	database.InitDB()
	migration.AutoMigrate()
	defer database.CloseDB()
	api.RunApp()
}
