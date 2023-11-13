package clients_test

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/oy1978/EdgeAPI/internal/db/models/clients"
)

func TestClientSystemDAO_CreateSystemIfNotExists(t *testing.T) {
	var dao = clients.NewClientSystemDAO()
	{
		err := dao.CreateSystemIfNotExists(nil, "Mac OS X")
		if err != nil {
			t.Fatal(err)
		}
	}
	{
		err := dao.CreateSystemIfNotExists(nil, "Mac OS X 2")
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestClientSystemDAO_Clean(t *testing.T) {
	var dao = clients.NewClientSystemDAO()
	err := dao.Clean(nil, 30)
	if err != nil {
		t.Fatal(err)
	}
}
