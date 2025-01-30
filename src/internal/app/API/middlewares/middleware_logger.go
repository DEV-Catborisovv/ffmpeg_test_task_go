package middlewares

import (
	"log"
	"os"
	"sync"
)

type LoggerMiddleWare struct {
	middleware Middleware
}

var loggerInstanceOnce sync.Once
var LoggerInstance *LoggerMiddleWare

func GetLoggerInstance() *LoggerMiddleWare {
	loggerInstanceOnce.Do(func() {
		LoggerInstance = &LoggerMiddleWare{}
	})
	return LoggerInstance
}

func NewLoggerMiddleware() *LoggerMiddleWare {
	return &LoggerMiddleWare{}
}

func (s *LoggerMiddleWare) Message(message string) {
	file, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Println("Error opening log file:", err)
		return
	}
	defer file.Close()

	logger := log.New(file, "", log.LstdFlags)

	logger.Println(message)
}
