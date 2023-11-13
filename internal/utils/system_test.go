// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package utils_test

import (
	"testing"

	"github.com/oy1978/EdgeAPI/internal/utils"
)

func TestSystemMemoryGB(t *testing.T) {
	t.Log(utils.SystemMemoryGB())
	t.Log(utils.SystemMemoryGB())
	t.Log(utils.SystemMemoryGB())
}
