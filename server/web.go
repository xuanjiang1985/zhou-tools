package server

import (
	"context"
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
	"zhou/tools/config"
	"zhou/tools/logger"

	"github.com/gin-gonic/gin"
)

var (
	//go:embed tmpl
	tmpl embed.FS

	//go:embed css js img font
	static embed.FS
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

	if "produce" == setting.AppEnv {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	// t, _ := template.ParseFS(public, "*.hmtl")
	// r.SetHTMLTemplate(t)
	// r.GET("/", func(ctx *gin.Context) {
	// 	ctx.HTML(200, "index.html", gin.H{"title": "Golang Embed 测试"})

	// })

	t, _ := template.ParseFS(tmpl, "tmpl/*.html")
	r.SetHTMLTemplate(t)
	version := time.Now().UnixNano()
	r.StaticFS("/static", http.FS(static))

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"version": version,
		})
	})

	r.GET("/zhou", func(c *gin.Context) {
		c.HTML(http.StatusOK, "zhou.html", gin.H{
			"version": version,
		})
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"title": "Main site",
		})
	})

	fmt.Printf("\n🚀 请访问网站: http://127.0.0.1:8900\n\n")
	err = r.Run(":8900")
	if err != nil {
		fmt.Printf("❌ 错误 %v\n", err)
		os.Exit(0)
	}
}
