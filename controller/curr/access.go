package curr

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Access struct {
	Url    string `json:"url"`
	Method string `json:"method"`
}

func (a *Access) HttpRewrite(c *gin.Context) {

	if a.Method == "GET" {
		m := make(map[string]interface{})
		m["method"] = a.Method

		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"content": m,
			"message": "ok",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"content": "",
		"message": "ok",
	})
}
