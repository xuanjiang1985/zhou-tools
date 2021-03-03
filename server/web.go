package server

import (
	"context"
	"fmt"
	"log"
	"time"
	"zhou/tools/config"
	"zhou/tools/logger"

	"github.com/gin-gonic/gin"
)

// StartWebServer start web server
func StartWebServer(ctx context.Context, ticker *time.Ticker) {
	// ticker
	go func() {
		for {
			select {
			case <-ctx.Done():
				logger.Logger.Info("web 服务器退出")
				return
			case <-ticker.C:
				logger.Logger.Info(("运行中"))
			}
		}
	}()

	setting, err := config.Read()
	if err != nil {
		log.Fatalln(err)
	}
	logger.Logger.Info(setting.AppEnv)
	if "produce" == setting.AppEnv {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	fmt.Printf("\n🚀 请访问网站: http://127.0.0.1:8900\n\n")
	r.Run(":8900")
}
