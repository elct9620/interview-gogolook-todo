package todo

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListTasks(c *gin.Context) {
	usecase := NewTodoUseCase()

	c.JSON(http.StatusOK, gin.H{
		"result": usecase.ListTasks(),
	})
}
