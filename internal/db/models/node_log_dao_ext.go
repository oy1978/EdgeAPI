// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .
//go:build !plus

package models

import (
	"github.com/iwind/TeaGo/dbs"
	"github.com/oy1978/EdgeCommon/pkg/nodeconfigs"
)

func (this *NodeLogDAO) deleteNodeLogsWithCluster(tx *dbs.Tx, role nodeconfigs.NodeRole, clusterId int64) error {
	return nil
}
