package regions

import (
	"encoding/json"

	"github.com/iwind/TeaGo/lists"
	"github.com/oy1978/EdgeAPI/internal/remotelogs"
)

func (this *RegionTown) DecodeCodes() []string {
	if len(this.Codes) == 0 {
		return []string{}
	}
	var result = []string{}
	err := json.Unmarshal(this.Codes, &result)
	if err != nil {
		remotelogs.Error("RegionTown.DecodeCodes", err.Error())
	}
	return result
}

func (this *RegionTown) DecodeCustomCodes() []string {
	if len(this.CustomCodes) == 0 {
		return []string{}
	}
	var result = []string{}
	err := json.Unmarshal(this.CustomCodes, &result)
	if err != nil {
		remotelogs.Error("RegionTown.DecodeCustomCodes", err.Error())
	}
	return result
}

func (this *RegionTown) DisplayName() string {
	if len(this.CustomName) > 0 {
		return this.CustomName
	}
	return this.Name
}

func (this *RegionTown) AllCodes() []string {
	var codes = this.DecodeCodes()

	if len(this.Name) > 0 && !lists.ContainsString(codes, this.Name) {
		codes = append(codes, this.Name)
	}

	if len(this.CustomName) > 0 && !lists.ContainsString(codes, this.CustomName) {
		codes = append(codes, this.CustomName)
	}

	for _, code := range this.DecodeCustomCodes() {
		if !lists.ContainsString(codes, code) {
			codes = append(codes, code)
		}
	}
	return codes
}
