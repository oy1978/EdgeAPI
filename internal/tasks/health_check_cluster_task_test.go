package tasks_test

import (
	"testing"

	"github.com/iwind/TeaGo/dbs"
	"github.com/oy1978/EdgeAPI/internal/tasks"
)

func TestHealthCheckClusterTask_Loop(t *testing.T) {
	dbs.NotifyReady()
	var task = tasks.NewHealthCheckClusterTask(10, nil)
	err := task.Loop()
	if err != nil {
		t.Fatal(err)
	}
	t.Log("ok")
}
