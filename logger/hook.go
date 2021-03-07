package logger

import (
	"fmt"
	"log"
	"os"
	"time"
	"zhou/tools/config"

	"github.com/sirupsen/logrus"
)

type logHook struct{}

func (h *logHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (h *logHook) Fire(entry *logrus.Entry) (err error) {
	fileName := time.Now().Format("2006-01-02")
	fullPath := fmt.Sprintf("%s/%s.log", config.Setting.Log.Logdir, fileName)

	if err = os.MkdirAll(config.Setting.Log.Logdir, os.ModePerm); err != nil {
		log.Panicln("创建文件夹错误", err)
		return
	}

	openFile, err := os.OpenFile(fullPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Panicln("写入日志文件错误", err)
		return
	}

	//设置输出
	Logger.Out = openFile
	return
}
