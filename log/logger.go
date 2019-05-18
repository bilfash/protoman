package log

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
)

func init() {
	fmt.Println("init logger called")
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func Get() *logrus.Entry {
	return logrus.WithField("application", os.Getenv("APP_NAME"))
}