package helpers_test

import (
	"github.com/iwind/TeaGo/Tea"
	_ "github.com/iwind/TeaGo/bootstrap"
	"github.com/oy1978/EdgeAPI/internal/installers/helpers"
	"testing"
)

func TestUnzip_Run(t *testing.T) {
	var unzip = helpers.NewUnzip(Tea.Root+"/deploy/edge-node-v0.0.1.zip", Tea.Root+"/deploy/")
	err := unzip.Run()
	if err != nil {
		t.Fatal(err)
	}
	t.Log("OK")
}
