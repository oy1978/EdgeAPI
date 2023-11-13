// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .

package utils_test

import (
	"testing"

	"github.com/iwind/TeaGo/types"
	"github.com/oy1978/EdgeAPI/internal/utils"
)

func TestSha1Random(t *testing.T) {
	for i := 0; i < 10; i++ {
		var s = utils.Sha1RandomString()
		t.Log("["+types.String(len(s))+"]", s)
	}
}
