package rest

import (
	"github.com/elct9620/interview-gogolook-todo/internal/todo"
	"github.com/gin-gonic/gin"
)

func AddTodoRoutes(router gin.IRouter, usecase *todo.TodoUseCase) {
	router.GET("/tasks", todo.ListTasks(usecase))
	router.POST("/tasks", todo.CreateTask(usecase))
	router.PUT("/tasks/:id", todo.UpdateTask(usecase))
	router.DELETE("/tasks/:id", todo.DeleteTask(usecase))
}
