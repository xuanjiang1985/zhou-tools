package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"
	"zhou/tools/logger"
	"zhou/tools/server"
)

func main() {

	// 初始化日志
	logger.Setup()

	ctx, cancel := context.WithCancel(context.Background())
	ticker := time.NewTicker(30 * time.Second)

	go server.StartWebServer(ctx, ticker)

	// 程序终止信号
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGHUP)

	// 退出信号
	s := <-signalChan
	cancel()
	logger.Logger.Info("所有进程退出:", s)
	time.Sleep(500 * time.Microsecond)
}
