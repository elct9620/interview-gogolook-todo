package rest

import (
	"github.com/elct9620/interview-gogolook-todo/internal/todo"
	"github.com/gin-gonic/gin"
)

func AddTodoRoutes(router gin.IRouter) {
	router.GET("/tasks", todo.ListTasks)
}
