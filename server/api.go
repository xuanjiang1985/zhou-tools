package server

import (
	"net/http"

	"zhou/tools/controller"
	"zhou/tools/controller/curr"
	"zhou/tools/controller/ws"

	"github.com/gin-gonic/gin"
)

// LoadRouterAPI contains api routes
func loadRouterAPI(e *gin.Engine) {
	router := e.Group("/api")
	{
		router.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"code":    0,
				"content": "Main site api",
				"message": "",
			})
		})

		router.POST("/scan/port", func(c *gin.Context) {
			port := &controller.Port{}
			c.BindJSON(port)

			port.Scan(c)
		})

		router.GET("/ws/echo", func(c *gin.Context) {
			ws.EchoMessage(c.Writer, c.Request)
		})

		router.POST("/curr/access", func(c *gin.Context) {
			access := &curr.Access{}
			c.BindJSON(access)
			access.HttpRewrite(c)
		})
	}
}
