package rest_test

import (
	"bytes"
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
	suite := NewTodoTestSuite(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/tasks", nil)
	suite.engine.ServeHTTP(res, req)

	if !cmp.Equal(http.StatusOK, res.Code) {
		t.Error(cmp.Diff(http.StatusOK, res.Code))
	}

	if !cmp.Equal(`{"result":[{"id":1,"name":"name","status":0}]}`, res.Body.String()) {
		t.Error(cmp.Diff(`{"result":[{"id":1,"name":"name","status":0}]}`, res.Body.String()))
	}
}

func Test_CreateTask(t *testing.T) {
	suite := NewTodoTestSuite(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/tasks", bytes.NewBufferString(`{"name":"買晚餐"}`))
	req.Header.Add("Content-Type", "application/json")
	suite.engine.ServeHTTP(res, req)

	if !cmp.Equal(http.StatusOK, res.Code) {
		t.Error(cmp.Diff(http.StatusOK, res.Code))
	}

	if !cmp.Equal(`{"result":{"id":1,"name":"買晚餐","status":0}}`, res.Body.String()) {
		t.Error(cmp.Diff(`{"result":{"id":1,"name":"買晚餐","status":0}}`, res.Body.String()))
	}
}
