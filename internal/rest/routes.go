package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddTodoRoutes(router gin.IRouter) {
	router.GET("/tasks", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"result": []gin.H{
				{"id": 1, "name": "name", "status": 0},
			},
		})
	})
}
