package telemetrygt

import (
	"log"
	"os"
)

var logger *log.Logger

func NewLogger() *log.Logger {
	if logger == nil {
		logger = log.New(os.Stdout, "", log.LstdFlags)
	}

	return logger
}

func Info(v ...interface{}) {
	logger.Println(v...)
}
