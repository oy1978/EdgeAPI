// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .

package installers_test

import (
	"github.com/oy1978/EdgeAPI/internal/installers"
	"github.com/oy1978/EdgeAPI/internal/utils/sizes"
	"github.com/oy1978/EdgeCommon/pkg/nodeconfigs"
	"testing"
	"time"
)

func TestNewUpgradeLimiter(t *testing.T) {
	var limiter = installers.NewUpgradeLimiter()
	limiter.UpdateNodeBytes(nodeconfigs.NodeRoleNode, 1, 1)
	limiter.UpdateNodeBytes(nodeconfigs.NodeRoleNode, 2, 5*sizes.M)
	t.Log("limiter:", limiter)
	t.Log("canUpgrade:", limiter.CanUpgrade())

	time.Sleep(1 * time.Second)
	t.Log("canUpgrade:", limiter.CanUpgrade())
	t.Log("limiter:", limiter)
	limiter.UpdateNodeBytes(nodeconfigs.NodeRoleNode, 2, 4*sizes.M)
	t.Log("canUpgrade:", limiter.CanUpgrade())

	t.Log("limiter:", limiter)
}
