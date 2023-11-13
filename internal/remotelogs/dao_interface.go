// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package remotelogs

import (
	"github.com/iwind/TeaGo/dbs"
	"github.com/oy1978/EdgeCommon/pkg/nodeconfigs"
)

type DAOInterface interface {
	CreateLog(tx *dbs.Tx, nodeRole nodeconfigs.NodeRole, nodeId int64, serverId int64, originId int64, level string, tag string, description string, createdAt int64, logType string, paramsJSON []byte) error
}
