package main

import (
	"github.com/elct9620/interview-gogolook-todo/internal/rest"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	rest.AddTodoRoutes(r)

	r.Run()
}
