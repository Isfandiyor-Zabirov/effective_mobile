package logger

import (
	"effective_mobile_tech_task/utils/env"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"log"
	"os"
	"time"
)

// SetLogger Установка Logger-а
var (
	Info  *log.Logger
	Error *log.Logger
)

func Init() {

	settings := env.GetSettings()

	path := "logs"
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}

	fileInfo, err := os.OpenFile(settings.Log.LogInfo, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
	}
	fileError, err := os.OpenFile(settings.Log.LogError, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
	}

	if err != nil {
		return
	}

	Info = log.New(fileInfo, "", log.Ldate|log.Lmicroseconds)
	Error = log.New(fileError, "", log.Ldate|log.Lmicroseconds)

	lumberLogInfo := &lumberjack.Logger{
		Filename:   settings.Log.LogInfo,
		MaxSize:    settings.Log.LogMaxSize, // megabytes
		MaxBackups: settings.Log.LogMaxBackups,
		MaxAge:     settings.Log.LogMaxAge,   //days
		Compress:   settings.Log.LogCompress, // disabled by default
		LocalTime:  true,
	}

	lumberLogError := &lumberjack.Logger{
		Filename:   settings.Log.LogError,
		MaxSize:    settings.Log.LogMaxSize, // megabytes
		MaxBackups: settings.Log.LogMaxBackups,
		MaxAge:     settings.Log.LogMaxAge,   //days
		Compress:   settings.Log.LogCompress, // disabled by default
		LocalTime:  true,
	}

	gin.DefaultWriter = io.MultiWriter(os.Stdout, lumberLogInfo)

	Info.SetOutput(gin.DefaultWriter)
	Error.SetOutput(lumberLogError)
}

// FormatLogs Форматирование логов
func FormatLogs(r *gin.Engine) {
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// your custom format
		return fmt.Sprintf("[GIN] %s - [%s] \"%s\" %s %s %d %s \"%s\" %s\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
}
