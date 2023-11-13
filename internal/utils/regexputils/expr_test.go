// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .

package regexputils_test

import (
	"testing"

	"github.com/iwind/TeaGo/assert"
	"github.com/oy1978/EdgeAPI/internal/utils/regexputils"
)

func TestExpr(t *testing.T) {
	var a = assert.NewAssertion(t)

	a.IsTrue(regexputils.YYYYMMDD.MatchString("20221011"))
	a.IsFalse(regexputils.YYYYMMDD.MatchString("202210"))

	a.IsTrue(regexputils.YYYYMM.MatchString("202210"))
	a.IsFalse(regexputils.YYYYMM.MatchString("20221011"))
}
