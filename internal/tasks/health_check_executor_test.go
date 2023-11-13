//go:build plus
// +build plus

package tasks_test

import (
	"testing"

	"github.com/iwind/TeaGo/dbs"
	teaconst "github.com/oy1978/EdgeAPI/internal/const"
	"github.com/oy1978/EdgeAPI/internal/tasks"
)

func TestHealthCheckExecutor_Run(t *testing.T) {
	teaconst.IsPlus = true
	dbs.NotifyReady()

	var executor = tasks.NewHealthCheckExecutor(42)
	results, err := executor.Run()
	if err != nil {
		t.Fatal(err)
	}
	for _, result := range results {
		t.Log(result.Node.Name, "addr:", result.NodeAddr, "isOk:", result.IsOk, "error:", result.Error)
	}
}
