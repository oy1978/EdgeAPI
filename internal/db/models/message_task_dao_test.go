package models_test

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/iwind/TeaGo/bootstrap"
	"github.com/iwind/TeaGo/dbs"
	"github.com/oy1978/EdgeAPI/internal/db/models"
)

func TestMessageTaskDAO_CleanExpiredMessageTasks(t *testing.T) {
	var dao = models.NewMessageTaskDAO()
	var tx *dbs.Tx
	err := dao.CleanExpiredMessageTasks(tx, 30)
	if err != nil {
		t.Fatal(err)
	}
}
