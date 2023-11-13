package models_test

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/iwind/TeaGo/bootstrap"
	"github.com/oy1978/EdgeAPI/internal/db/models"
)

func TestMetricStatDAO_Clean(t *testing.T) {
	var dao = models.NewMetricStatDAO()
	t.Log(dao.Clean(nil))
}
