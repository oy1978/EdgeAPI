package models_test

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/iwind/TeaGo/bootstrap"
	"github.com/iwind/TeaGo/dbs"
	"github.com/oy1978/EdgeAPI/internal/db/models"
)

func TestHTTPCacheTaskDAO_Clean(t *testing.T) {
	dbs.NotifyReady()

	err := models.SharedHTTPCacheTaskDAO.CleanDays(nil, 30)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("ok")
}
