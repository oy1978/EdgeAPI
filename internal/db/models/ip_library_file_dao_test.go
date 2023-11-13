package models_test

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/iwind/TeaGo/bootstrap"
	"github.com/iwind/TeaGo/dbs"
	"github.com/oy1978/EdgeAPI/internal/db/models"
)

func TestIPLibraryFileDAO_GenerateIPLibrary(t *testing.T) {
	dbs.NotifyReady()

	var tx *dbs.Tx
	err := models.SharedIPLibraryFileDAO.GenerateIPLibrary(tx, 4)
	if err != nil {
		t.Fatal(err)
	}
}
