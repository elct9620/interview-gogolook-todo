package rest

import (
	"github.com/elct9620/interview-gogolook-todo/internal/todo"
	"github.com/elct9620/interview-gogolook-todo/internal/todo/repo"
	"github.com/gin-gonic/gin"
)

func AddTodoRoutes(router gin.IRouter) {
	repo := repo.NewMemoryStore()
	usecase := todo.NewTodoUseCase(repo)

	router.GET("/tasks", todo.ListTasks(usecase))
	router.POST("/tasks", todo.CreateTask(usecase))
}
