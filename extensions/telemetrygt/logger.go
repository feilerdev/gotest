package telemetrygt

import (
	"log"
	"os"
)

var logger *log.Logger

// TODO(name.surname): Add test -> td-quality | Important.

func NewLogger() *log.Logger {
	if logger == nil {
		logger = log.New(os.Stdout, "", log.LstdFlags)
	}

	logger.Println("created logger successfully")

	return logger
}

func Info(v ...interface{}) {
	logger.Println(v...)
}
