package tasks_test

import (
	"testing"
	"time"

	"github.com/iwind/TeaGo/dbs"
	"github.com/oy1978/EdgeAPI/internal/tasks"
)

func TestLogTask_LoopClean(t *testing.T) {
	dbs.NotifyReady()

	var task = tasks.NewLogTask(24*time.Hour, 1*time.Minute)
	err := task.LoopClean()
	if err != nil {
		t.Fatal(err)
	}
	t.Log("ok")
}

func TestLogTask_LoopMonitor(t *testing.T) {
	dbs.NotifyReady()

	var task = tasks.NewLogTask(24*time.Hour, 1*time.Minute)
	err := task.LoopMonitor()
	if err != nil {
		t.Fatal(err)
	}
	t.Log("ok")
}
