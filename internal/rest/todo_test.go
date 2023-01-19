package rest_test

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/elct9620/interview-gogolook-todo/internal/rest"
	"github.com/elct9620/interview-gogolook-todo/internal/todo"
	"github.com/elct9620/interview-gogolook-todo/internal/todo/repo"
	"github.com/gin-gonic/gin"
	"github.com/google/go-cmp/cmp"
)

type TodoTestSuite struct {
	engine *gin.Engine
	repo   todo.TodoRepository
}

func NewTodoTestSuite(t *testing.T) *TodoTestSuite {
	repo := repo.NewMemoryStore()
	usecase := todo.NewTodoUseCase(repo)

	engine := gin.Default()
	rest.AddTodoRoutes(engine, usecase)

	return &TodoTestSuite{
		engine: engine,
		repo:   repo,
	}
}

func Test_GetTasks(t *testing.T) {
	tests := []struct {
		tasks    []string
		expected string
	}{
		{
			tasks:    []string{"name"},
			expected: `{"result":[{"id":1,"name":"name","status":0}]}`,
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(fmt.Sprintf("test %s", tc.expected), func(tt *testing.T) {
			tt.Parallel()

			suite := NewTodoTestSuite(t)

			for _, task := range tc.tasks {
				suite.repo.CreateTask(task)
			}

			res := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/tasks", nil)
			suite.engine.ServeHTTP(res, req)

			if !cmp.Equal(http.StatusOK, res.Code) {
				tt.Error(cmp.Diff(http.StatusOK, res.Code))
			}

			if !cmp.Equal(tc.expected, res.Body.String()) {
				tt.Error(cmp.Diff(tc.expected, res.Body.String()))
			}
		})
	}
}

func Test_CreateTask(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    `{"name":"買晚餐"}`,
			expected: `{"result":{"id":1,"name":"買晚餐","status":0}}`,
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(fmt.Sprintf("test %s", tc.input), func(tt *testing.T) {
			tt.Parallel()

			suite := NewTodoTestSuite(t)

			res := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/tasks", bytes.NewBufferString(tc.input))
			req.Header.Add("Content-Type", "application/json")
			suite.engine.ServeHTTP(res, req)

			if !cmp.Equal(http.StatusOK, res.Code) {
				tt.Error(cmp.Diff(http.StatusOK, res.Code))
			}

			if !cmp.Equal(tc.expected, res.Body.String()) {
				tt.Error(cmp.Diff(tc.expected, res.Body.String()))
			}
		})
	}
}
