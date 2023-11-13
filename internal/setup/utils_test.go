// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package setup_test

import (
	"testing"

	"github.com/oy1978/EdgeAPI/internal/setup"
)

func TestComposeSQLVersion(t *testing.T) {
	t.Log(setup.ComposeSQLVersion())
}
