package log

import (
	"github.com/bilfash/protoman/config"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func Get() *logrus.Entry {
	return logrus.WithField("application", config.AppName())
}