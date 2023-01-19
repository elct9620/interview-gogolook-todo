package rest_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/elct9620/interview-gogolook-todo/internal/rest"
	"github.com/gin-gonic/gin"
	"github.com/google/go-cmp/cmp"
)

type TodoTestSuite struct {
	engine *gin.Engine
}

func NewTodoTestSuite(t *testing.T) *TodoTestSuite {
	engine := gin.Default()
	rest.AddTodoRoutes(engine)

	return &TodoTestSuite{
		engine: engine,
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
	req := httptest.NewRequest(http.MethodPost, "/tasks", nil)
	suite.engine.ServeHTTP(res, req)

	if !cmp.Equal(http.StatusOK, res.Code) {
		t.Error(cmp.Diff(http.StatusOK, res.Code))
	}

	if !cmp.Equal(`{"result":{"id":1,"name":"買晚餐","status":0}}`, res.Body.String()) {
		t.Error(cmp.Diff(`{"result":{"id":1,"name":"買晚餐","status":0}}`, res.Body.String()))
	}
}
