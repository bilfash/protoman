package log

import (
	"github.com/sirupsen/logrus"
	"os"
)

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func Get() *logrus.Entry {
	return logrus.WithField("application", os.Getenv("APP_NAME"))
}