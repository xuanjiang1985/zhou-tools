package logger

import (
	"log"
	"zhou/tools/config"

	"github.com/sirupsen/logrus"
)

// Logger is a new instance
var Logger = logrus.New()

// Setup init logrus
func Setup() {

	level, err := logrus.ParseLevel(config.Setting.Log.Level)
	if err != nil {
		log.Panicln("日志level格式设置错误", err)
	}
	Logger.SetLevel(level)

	Logger.SetFormatter(&logrus.JSONFormatter{})

	// 取消线程安全
	Logger.SetNoLock()

	// 自定义HOOK
	Logger.AddHook(&logHook{})
}
