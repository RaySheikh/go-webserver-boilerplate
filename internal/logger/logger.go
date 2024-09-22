package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

// SetupLogger sets up structured logging using logrus
func SetupLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(os.Stdout)
	logger.SetLevel(logrus.InfoLevel)
	return logger
}
