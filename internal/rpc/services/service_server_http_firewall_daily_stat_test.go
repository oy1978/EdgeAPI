package services

import (
	"github.com/iwind/TeaGo/dbs"
	"github.com/iwind/TeaGo/logs"
	rpcutils "github.com/oy1978/EdgeAPI/internal/rpc/utils"
	"github.com/oy1978/EdgeCommon/pkg/rpc/pb"
	"testing"
)

func TestServerHTTPFirewallDailyStatService_ComposeServerHTTPFirewallDashboard(t *testing.T) {
	dbs.NotifyReady()

	service := new(ServerHTTPFirewallDailyStatService)
	resp, err := service.ComposeServerHTTPFirewallDashboard(rpcutils.NewMockAdminNodeContext(1), &pb.ComposeServerHTTPFirewallDashboardRequest{})
	if err != nil {
		t.Fatal(err)
	}
	logs.PrintAsJSON(resp, t)
}
