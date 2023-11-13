package services

import (
	"github.com/iwind/TeaGo/dbs"
	"github.com/iwind/TeaGo/logs"
	rpcutils "github.com/oy1978/EdgeAPI/internal/rpc/utils"
	"github.com/oy1978/EdgeCommon/pkg/rpc/pb"
	"testing"
)

func TestHTTPFirewallPolicyService_CheckHTTPFirewallPolicyIPStatus(t *testing.T) {
	dbs.NotifyReady()
	service := &HTTPFirewallPolicyService{}

	{
		resp, err := service.CheckHTTPFirewallPolicyIPStatus(rpcutils.NewMockAdminNodeContext(1), &pb.CheckHTTPFirewallPolicyIPStatusRequest{
			HttpFirewallPolicyId: 14,
			Ip:                   "127.0.0.1",
		})
		if err != nil {
			t.Fatal(err)
		}
		logs.PrintAsJSON(resp, t)
	}

	{
		resp, err := service.CheckHTTPFirewallPolicyIPStatus(rpcutils.NewMockAdminNodeContext(1), &pb.CheckHTTPFirewallPolicyIPStatusRequest{
			HttpFirewallPolicyId: 14,
			Ip:                   "192.168.1.100",
		})
		if err != nil {
			t.Fatal(err)
		}
		logs.PrintAsJSON(resp, t)
	}

	{
		resp, err := service.CheckHTTPFirewallPolicyIPStatus(rpcutils.NewMockAdminNodeContext(1), &pb.CheckHTTPFirewallPolicyIPStatusRequest{
			HttpFirewallPolicyId: 14,
			Ip:                   "221.218.201.94",
		})
		if err != nil {
			t.Fatal(err)
		}
		logs.PrintAsJSON(resp, t)
	}
}
