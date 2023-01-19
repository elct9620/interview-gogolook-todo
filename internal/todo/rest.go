package todo

import (
	"net/http"

	"github.com/elct9620/interview-gogolook-todo/internal/todo/repo"
	"github.com/gin-gonic/gin"
)

func ListTasks(c *gin.Context) {
	repo := repo.NewMemoryStore()
	usecase := NewTodoUseCase(repo)

	c.JSON(http.StatusOK, gin.H{
		"result": usecase.ListTasks(),
	})
}

func CreateTask(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"result": gin.H{"id": 1, "name": "買晚餐", "status": 0},
	})
}
