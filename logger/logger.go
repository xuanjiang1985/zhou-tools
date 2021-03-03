package logger

import (
	"log"
	"os"

	"github.com/sirupsen/logrus"
)

// Logger is a new instance
var Logger = logrus.New()

func init() {

	src, err := os.OpenFile("./logrus.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("create file log failed: %v", err)
	}
	Logger.Out = src
	Logger.SetFormatter(&logrus.JSONFormatter{})
}
