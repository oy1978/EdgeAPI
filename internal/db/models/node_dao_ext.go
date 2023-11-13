// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved.
//go:build !plus
// +build !plus

package models

import (
	"github.com/iwind/TeaGo/dbs"
	"github.com/oy1978/EdgeAPI/internal/utils"
	"github.com/oy1978/EdgeAPI/internal/zero"
	"github.com/oy1978/EdgeCommon/pkg/nodeconfigs"
)

func (this *NodeDAO) loadServersFromCluster(tx *dbs.Tx, clusterId int64, serverIdMap map[int64]zero.Zero) ([]*Server, error) {
	return nil, nil
}

func (this *NodeDAO) composeExtConfig(tx *dbs.Tx, config *nodeconfigs.NodeConfig, clusterIds []int64, cacheMap *utils.CacheMap) error {
	return nil
}

// CheckNodeIPAddresses 检查节点IP地址
func (this *NodeDAO) CheckNodeIPAddresses(tx *dbs.Tx, node *Node) (shouldSkip bool, shouldOverwrite bool, ipAddressStrings []string, err error) {
	return
}
