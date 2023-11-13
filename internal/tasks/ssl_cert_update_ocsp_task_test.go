// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package tasks_test

import (
	"testing"
	"time"

	"github.com/iwind/TeaGo/dbs"
	"github.com/oy1978/EdgeAPI/internal/tasks"
)

func TestSSLCertUpdateOCSPTask_Loop(t *testing.T) {
	dbs.NotifyReady()

	var task = tasks.NewSSLCertUpdateOCSPTask(1 * time.Minute)
	err := task.Loop()
	if err != nil {
		t.Fatal(err)
	}
}
