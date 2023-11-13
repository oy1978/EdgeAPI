// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .
//go:build !plus

package services

import (
	"context"
	"github.com/iwind/TeaGo/dbs"
	"github.com/oy1978/EdgeCommon/pkg/rpc/pb"
)

// ComposeAdminDashboard方法扩展
func (this *AdminService) composeAdminDashboardExt(tx *dbs.Tx, ctx context.Context, result *pb.ComposeAdminDashboardResponse) error {
	return nil
}
