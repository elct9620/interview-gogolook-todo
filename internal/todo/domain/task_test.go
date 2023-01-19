package domain_test

import (
	"testing"

	"github.com/elct9620/interview-gogolook-todo/internal/todo/domain"
	"github.com/google/go-cmp/cmp"
)

func Test_Task_Rename(t *testing.T) {
	task := domain.NewTask(1, "買晚餐", false)
	task.Rename("買早餐")

	if !cmp.Equal("買早餐", task.Name) {
		t.Error(cmp.Diff("買早餐", task.Name))
	}
}

func Test_Task_MarkComplete(t *testing.T) {
	task := domain.NewTask(1, "買晚餐", false)
	task.MarkComplete()

	if !cmp.Equal(true, task.Completed) {
		t.Error(cmp.Diff(true, task.Completed))
	}

	if !cmp.Equal(1, task.StatusCode()) {
		t.Error(cmp.Diff(1, task.StatusCode()))
	}
}

func Test_Task_MarkIncomplete(t *testing.T) {
	task := domain.NewTask(1, "買晚餐", true)
	task.MarkIncomplete()

	if !cmp.Equal(false, task.Completed) {
		t.Error(cmp.Diff(false, task.Completed))
	}

	if !cmp.Equal(0, task.StatusCode()) {
		t.Error(cmp.Diff(0, task.StatusCode()))
	}
}
