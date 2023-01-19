package main

import (
	"github.com/elct9620/interview-gogolook-todo/internal/rest"
	"github.com/elct9620/interview-gogolook-todo/internal/todo"
	"github.com/elct9620/interview-gogolook-todo/internal/todo/repo"
	"github.com/gin-gonic/gin"
)

func main() {
	repo := repo.NewMemoryStore()
	usecase := todo.NewTodoUseCase(repo)

	r := gin.Default()
	rest.AddTodoRoutes(r, usecase)

	r.Run()
}
