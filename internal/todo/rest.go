package todo

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListTasks(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"result": []gin.H{
			{"id": 1, "name": "name", "status": 0},
		},
	})
}
