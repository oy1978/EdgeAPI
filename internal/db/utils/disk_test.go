// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .

package dbutils_test

import (
	_ "github.com/iwind/TeaGo/bootstrap"
	dbutils "github.com/oy1978/EdgeAPI/internal/db/utils"
	"testing"
)

func TestHasFreeSpace(t *testing.T) {
	t.Log(dbutils.CheckHasFreeSpace())
	t.Log(dbutils.LocalDatabaseDataDir)
}
