package tasks_test

import (
	"testing"
	"time"

	"github.com/iwind/TeaGo/dbs"
	timeutil "github.com/iwind/TeaGo/utils/time"
	"github.com/oy1978/EdgeAPI/internal/tasks"
)

func TestSSLCertExpireCheckExecutor_loop(t *testing.T) {
	dbs.NotifyReady()

	t.Log("30 days later: ", timeutil.FormatTime("Y-m-d", time.Now().Unix()+30*86400), time.Now().Unix()+30*86400)
	t.Log("14 days later: ", timeutil.FormatTime("Y-m-d", time.Now().Unix()+14*86400), time.Now().Unix()+14*86400)
	t.Log("7 days later: ", timeutil.FormatTime("Y-m-d", time.Now().Unix()+7*86400), time.Now().Unix()+7*86400)
	t.Log("3 days later: ", timeutil.FormatTime("Y-m-d", time.Now().Unix()+3*86400), time.Now().Unix()+3*86400)
	t.Log("today: ", timeutil.FormatTime("Y-m-d", time.Now().Unix()), time.Now().Unix())

	var task = tasks.NewSSLCertExpireCheckExecutor(1 * time.Hour)
	err := task.Loop()
	if err != nil {
		t.Fatal(err)
	}
	t.Log("ok")
}
