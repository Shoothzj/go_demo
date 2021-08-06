package demo_hc

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
	iotda "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/iotda/v5"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/iotda/v5/region"
	"go_demo/component_hc_config"
)

func getIotda() *iotda.IoTDAClient {
	hcConfig := component_hc_config.ReadHcConfig()
	auth := basic.NewCredentialsBuilder().WithAk(hcConfig.Ak).WithSk(hcConfig.Sk).Build()
	return iotda.NewIoTDAClient(iotda.IoTDAClientBuilder().WithRegion(region.CN_NORTH_4).WithCredential(auth).Build())
}
