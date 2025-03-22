package logger

import (
	"github.com/sirupsen/logrus"
)

var Logger = logrus.New()

func init() {
	Logger.SetFormatter(&logrus.JSONFormatter{})	
}

func Log(level logrus.Level, message string) {
	Logger.WithFields(logrus.Fields{
		"level": level,
		"message": message,

	}).Log(level, message)
}

func Info(message string) {
	Log(logrus.InfoLevel, message)
}

func Error(message string) {
	Log(logrus.ErrorLevel, message)
}	

func Debug(message string) {
	Log(logrus.DebugLevel, message)
}

func Warn(message string) {
	Log(logrus.WarnLevel, message)
}

func Fatal(message string) {
	Log(logrus.FatalLevel, message)
}