package models

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/iwind/TeaGo/bootstrap"
	"github.com/oy1978/EdgeCommon/pkg/nodeconfigs"
)

func TestNodeIPAddressDAO_FindFirstNodeAccessIPAddress(t *testing.T) {
	var dao = NewNodeIPAddressDAO()
	t.Log(dao.FindFirstNodeAccessIPAddress(nil, 48, true, nodeconfigs.NodeRoleNode))
	t.Log(dao.FindFirstNodeAccessIPAddressId(nil, 48, true, nodeconfigs.NodeRoleNode))
}
