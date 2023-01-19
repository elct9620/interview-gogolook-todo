package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/tasks", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"result": []gin.H{
				{"id": 1, "name": "name", "status": 0},
			},
		})
	})

	r.Run()
}
