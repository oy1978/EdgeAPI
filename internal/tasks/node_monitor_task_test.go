package tasks_test

import (
	"testing"
	"time"

	"github.com/iwind/TeaGo/dbs"
	"github.com/oy1978/EdgeAPI/internal/db/models"
	"github.com/oy1978/EdgeAPI/internal/tasks"
)

func TestNodeMonitorTask_loop(t *testing.T) {
	dbs.NotifyReady()

	var task = tasks.NewNodeMonitorTask(60 * time.Second)
	err := task.Loop()
	if err != nil {
		t.Fatal(err)
	}
	t.Log("ok")
}

func TestNodeMonitorTask_Monitor(t *testing.T) {
	dbs.NotifyReady()
	var task = tasks.NewNodeMonitorTask(60 * time.Second)
	for i := 0; i < 5; i++ {
		err := task.MonitorCluster(&models.NodeCluster{
			Id: 42,
		})
		if err != nil {
			t.Fatal(err)
		}
	}
}
