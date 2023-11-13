package models

import (
	"encoding/json"

	"github.com/oy1978/EdgeAPI/internal/remotelogs"
	"github.com/oy1978/EdgeCommon/pkg/serverconfigs/sslconfigs"
)

func (this *SSLPolicy) DecodeCerts() []*sslconfigs.SSLCertRef {
	if len(this.Certs) == 0 {
		return nil
	}

	var refs = []*sslconfigs.SSLCertRef{}
	err := json.Unmarshal(this.Certs, &refs)
	if err != nil {
		remotelogs.Error("SSLPolicy_DecodeCerts", err.Error())
	}
	return refs
}
