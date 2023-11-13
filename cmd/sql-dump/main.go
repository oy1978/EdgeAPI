package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/iwind/TeaGo/Tea"
	_ "github.com/iwind/TeaGo/bootstrap"
	"github.com/iwind/TeaGo/dbs"
	"github.com/oy1978/EdgeAPI/internal/setup"
)

func main() {
	db, err := dbs.Default()
	if err != nil {
		fmt.Println("[ERROR]" + err.Error())
		return
	}
	results, err := setup.NewSQLDump().Dump(db, true)
	if err != nil {
		fmt.Println("[ERROR]" + err.Error())
		return
	}

	prettyResultsJSON, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		fmt.Println("[ERROR]" + err.Error())
		return
	}

	// 写入到 sql.json 中
	var dir = filepath.Dir(Tea.Root)
	err = os.WriteFile(dir+"/internal/setup/sql.json", prettyResultsJSON, 0666)
	if err != nil {
		fmt.Println("[ERROR]" + err.Error())
		return
	}

	fmt.Println("ok")
}
