// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.
//go:build !plus
// +build !plus

package models

import (
	"github.com/iwind/TeaGo/dbs"
	"github.com/oy1978/EdgeCommon/pkg/nodeconfigs"
)

// FireThresholds 触发阈值
func (this *NodeIPAddressDAO) FireThresholds(tx *dbs.Tx, role nodeconfigs.NodeRole, nodeId int64) error {
	return nil
}
