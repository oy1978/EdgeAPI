// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .

package utils

import (
	"os/exec"
	"runtime"

	"github.com/iwind/TeaGo/types"
	"github.com/oy1978/EdgeAPI/internal/remotelogs"
	executils "github.com/oy1978/EdgeAPI/internal/utils/exec"
)

func AddPortsToFirewall(ports []int) {
	for _, port := range ports {
		// Linux
		if runtime.GOOS == "linux" {
			// firewalld
			firewallCmd, _ := executils.LookPath("firewall-cmd")
			if len(firewallCmd) > 0 {
				err := exec.Command(firewallCmd, "--add-port="+types.String(port)+"/tcp").Run()
				if err == nil {
					remotelogs.Println("API_NODE", "add port '"+types.String(port)+"' to firewalld")

					_ = exec.Command(firewallCmd, "--add-port="+types.String(port)+"/tcp", "--permanent").Run()
				}
			}
		}
	}
}
