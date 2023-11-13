package dns_test

import (
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/iwind/TeaGo/bootstrap"
	"github.com/iwind/TeaGo/dbs"
	"github.com/oy1978/EdgeAPI/internal/db/models/dns"
)

func TestDNSTaskDAO_CreateDNSTask(t *testing.T) {
	dbs.NotifyReady()
	err := dns.SharedDNSTaskDAO.CreateDNSTask(nil, 1, 2, 3, 0, "cdn", "taskType")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("ok")
}

func TestDNSTaskDAO_UpdateClusterDNSTasksDone(t *testing.T) {
	var dao = dns.NewDNSTaskDAO()
	var tx *dbs.Tx
	err := dao.UpdateClusterDNSTasksDone(tx, 46, time.Now().UnixNano())
	if err != nil {
		t.Fatal(err)
	}
}
