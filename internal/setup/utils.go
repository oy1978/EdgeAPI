// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package setup

import (
	"strings"

	teaconst "github.com/oy1978/EdgeAPI/internal/const"
)

func ComposeSQLVersion() string {
	var version = teaconst.Version
	if len(teaconst.SQLVersion) == 0 {
		return version
	}

	if strings.Count(version, ".") <= 2 {
		return version + "." + teaconst.SQLVersion
	}
	return version
}
