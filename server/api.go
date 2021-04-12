package server

import (
	"net/http"

	"zhou/tools/controller"

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
	}
}
