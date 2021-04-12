package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
	"zhou/tools/config"
	"zhou/tools/cron"
	"zhou/tools/logger"
	"zhou/tools/server"
)

func main() {

	// 读取配置
	if err := config.Setup(); err != nil {
		fmt.Println(err)
		return
	}

	//fmt.Println(config.Setting)

	// 初始化日志
	if err := logger.Setup(); err != nil {
		fmt.Println(err)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	ticker := time.NewTicker(60 * time.Second)

	// 任务调度
	go cron.Run(ctx)

	// web 服务
	go server.StartWebServer(ctx, ticker)

	// web socket

	// 程序终止信号
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL,
		syscall.SIGQUIT, syscall.SIGHUP, syscall.SIGKILL)

	// 退出信号
	s := <-signalChan
	cancel()
	logger.Logger.Info("捕捉到中断信号: ", s)
	time.Sleep(500 * time.Microsecond)
}
