package todo

import (
	"net/http"
	"strconv"

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
		var input CreateTaskInput
		c.ShouldBind(&input)

		c.JSON(http.StatusOK, gin.H{
			"result": usecase.CreateTask(&input),
		})
	}
}

func UpdateTask(usecase *TodoUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input UpdateTaskInput
		c.ShouldBind(&input)

		idStr := c.Param("id")
		id, _ := strconv.Atoi(idStr)

		c.JSON(http.StatusOK, gin.H{
			"result": usecase.UpdateTask(id, &input),
		})
	}
}
