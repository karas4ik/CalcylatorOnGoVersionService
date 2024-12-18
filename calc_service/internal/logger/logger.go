package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

func InitLogger(logLevel string) {
	log = logrus.New()
	level, err := logrus.ParseLevel(logLevel)
	if err == nil {
		log.SetLevel(level)
	} else {
		log.SetLevel(logrus.InfoLevel)
	}
	log.SetFormatter(&logrus.JSONFormatter{
		PrettyPrint: true,
	})
	file, err := os.OpenFile("calc_service.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.SetOutput(file)
	}
}

func GetLogger() *logrus.Logger {
	return log
}
