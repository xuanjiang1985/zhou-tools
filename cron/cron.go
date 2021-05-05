package cron

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"
	"zhou/tools/config"

	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
)

// 文件日志
type logHook struct{}

func (h *logHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (h *logHook) Fire(entry *logrus.Entry) (err error) {
	fileName := time.Now().Format("2006-01-02")
	fullPath := fmt.Sprintf("%s/%s.cron.log", config.Setting.Log.Logdir, fileName)

	if err = os.MkdirAll(config.Setting.Log.Logdir, os.ModePerm); err != nil {
		log.Panic("创建文件夹错误", err)
		return
	}

	openFile, err := os.OpenFile(fullPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Panic("写入日志文件错误", err)
		return
	}

	//设置输出
	logger.Out = openFile
	return
}

// # 文件格式說明
// # ┌──分鐘（0 - 59）
// # │  ┌──小時（0 - 23）
// # │  │  ┌──日（1 - 31）
// # │  │  │  ┌─月（1 - 12）
// # │  │  │  │  ┌─星期（0 - 6，表示从周日到周六）
// # │  │  │  │  │
// # *  *  *  *  * 被執行的命令

var logger = logrus.New()

// Run starts cron job
func Run(ctx context.Context) {

	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.AddHook(&logHook{})

	c := cron.New()

	// this task is running for test
	// i := 0
	// c.AddFunc("* * * * *", func() {
	// 	logger.WithFields(logrus.Fields{"i": i}).Info("每1分钟执行")
	// 	i++
	// })

	c.AddFunc("*/3 * * * *", func() {
		logger.WithField("cron", "running").Info("每3分钟执行")

	})

	// c.AddFunc("10 * * * *", func() {
	// 	logger.WithField("db_backup", 1).Info("每1小时执行")
	// 	cmd := exec.Command("/bin/bash", "-c", "/Users/zhougang/db_backup/myj_wyb.sh")
	// 	output, err := cmd.Output()
	// 	if err != nil {
	// 		logger.WithField("db_backup", 0).Info(err.Error())
	// 		return
	// 	}
	// 	logger.WithField("db_backup", 1).Info(string(output))
	// })

	// c.AddFunc("12 * * * *", func() {
	// 	logger.WithField("db_backup", 1).Info("每1小时执行")
	// 	cmd := exec.Command("/bin/bash", "-c", "/Users/zhougang/db_backup/hios_order.sh")
	// 	output, err := cmd.Output()
	// 	if err != nil {
	// 		logger.WithField("db_backup", 0).Info(err.Error())
	// 		return
	// 	}
	// 	logger.WithField("db_backup", 1).Info(string(output))
	// })

	// 启动执行任务
	c.Start()
	// 退出时关闭计划任务
	defer c.Stop()

	// 如果使用 select{} 那么就一直会循环
	select {
	case <-ctx.Done():
		logger.Info("cron job 停止")
		return
	}
}
