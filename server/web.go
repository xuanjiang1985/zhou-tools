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
	port = "8905"
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

	if "prod" == setting.AppEnv {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	r := gin.Default()

	t, _ := template.ParseFS(tmpl, "tmpl/*.html", "tmpl/**/*.html")
	r.SetHTMLTemplate(t)
	version := time.Now().UnixNano()
	r.StaticFS("/static", http.FS(static))

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title":   "工具管理",
			"version": version,
		})
	})

	r.GET("/zhou", func(c *gin.Context) {
		c.HTML(http.StatusOK, "zhou.html", gin.H{
			"title":   "工具管理",
			"version": version,
		})
	})

	r.GET("favicon.ico", func(c *gin.Context) {
		file, err := static.ReadFile("img/favicon.ico")
		if err != nil {
			logger.Logger.Info(err)
		}

		c.Data(
			http.StatusOK,
			"image/x-icon",
			file,
		)
	})

	loadRouterAPI(r)

	fmt.Printf("\n🚀 请访问网站: http://127.0.0.1:%s\n\n", port)
	err = r.Run(":" + port)
	if err != nil {
		fmt.Printf("❌ 错误 %v\n", err)
		os.Exit(0)
	}
}
