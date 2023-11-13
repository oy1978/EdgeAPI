package models

import (
	"encoding/json"

	"github.com/oy1978/EdgeAPI/internal/remotelogs"
	"github.com/oy1978/EdgeCommon/pkg/nodeconfigs"
)

func (this *NodeIPAddressThreshold) DecodeItems() (result []*nodeconfigs.IPAddressThresholdItemConfig) {
	if len(this.Items) == 0 {
		return
	}

	err := json.Unmarshal(this.Items, &result)
	if err != nil {
		remotelogs.Error("NodeIPAddressThreshold", "decode items: "+err.Error())
	}
	return
}

func (this *NodeIPAddressThreshold) DecodeActions() (result []*nodeconfigs.IPAddressThresholdActionConfig) {
	if len(this.Actions) == 0 {
		return
	}

	err := json.Unmarshal(this.Actions, &result)
	if err != nil {
		remotelogs.Error("NodeIPAddressThreshold", "decode actions: "+err.Error())
	}
	return
}
