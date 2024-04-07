package env

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

var settings Settings

func ReadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	settings.Database.Name = os.Getenv("DATABASE")
	settings.Database.Port = os.Getenv("DATABASE_PORT")
	settings.Database.Host = os.Getenv("DATABASE_HOST")
	settings.Database.Username = os.Getenv("DATABASE_USERNAME")
	settings.Database.Password = os.Getenv("DATABASE_PASSWORD")

	settings.Log.LogInfo = os.Getenv("INFO_LOG")
	settings.Log.LogDebug = os.Getenv("DEBUG_LOG")
	settings.Log.LogError = os.Getenv("ERROR_LOG")
	logMaxSize := os.Getenv("LOG_MAX_SIZE")
	logMaxBackups := os.Getenv("LOG_MAX_BACKUPS")
	logMaxAge := os.Getenv("LOG_MAX_AGE")
	logCompress := os.Getenv("LOG_COMPRESS")
	settings.Log.LogMaxSize, _ = strconv.Atoi(logMaxSize)
	settings.Log.LogMaxBackups, _ = strconv.Atoi(logMaxBackups)
	settings.Log.LogMaxAge, _ = strconv.Atoi(logMaxAge)
	settings.Log.LogCompress, _ = strconv.ParseBool(logCompress)

	settings.AppParams.Host = os.Getenv("APP_HOST")
	settings.AppParams.Port = os.Getenv("APP_PORT")

	settings.ExternalApi.GetCarInfoUrl = os.Getenv("GET_CARS_HOST")

	fmt.Println("settings:", settings)
}

func GetSettings() *Settings {
	return &settings
}
