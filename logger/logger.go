package logger

import (
	"zhou/tools/config"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// Logger is a new instance
var Logger = logrus.New()

// Setup init logrus
func Setup() error {

	level, err := logrus.ParseLevel(config.Setting.Log.Level)
	if err != nil {
		return errors.WithStack(err)
	}
	Logger.SetLevel(level)

	Logger.SetFormatter(&logrus.JSONFormatter{})

	// 取消线程安全
	Logger.SetNoLock()

	// 自定义HOOK
	Logger.AddHook(&logHook{})

	return nil
}
