package nameservers

import (
	"encoding/json"

	"github.com/oy1978/EdgeAPI/internal/db/models"
	"github.com/oy1978/EdgeAPI/internal/remotelogs"
)

func (this *NSDomain) DecodeGroupIds() []int64 {
	if models.IsNull(this.GroupIds) {
		return nil
	}

	var result = []int64{}
	err := json.Unmarshal(this.GroupIds, &result)
	if err != nil {
		remotelogs.Error("NSDomain", "DecodeGroupIds:"+err.Error())
	}
	return result
}
