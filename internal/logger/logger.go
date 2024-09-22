package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var (
	Logger *logrus.Logger
)

func InitLogger() {
	Logger = logrus.New()
	Logger.SetOutput(os.Stdout)
	Logger.SetFormatter(&logrus.JSONFormatter{}) // Use JSON format for structured logging
	Logger.SetLevel(logrus.InfoLevel)
}
