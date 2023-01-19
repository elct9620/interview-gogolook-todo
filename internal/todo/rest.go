package todo

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListTasks(usecase *TodoUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"result": usecase.ListTasks(),
		})
	}
}

func CreateTask(usecase *TodoUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"result": usecase.CreateTask("買晚餐"),
		})
	}
}
