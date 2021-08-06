package demo_hc

import (
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/config"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/live/v2/region"
	vpc "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vpc/v2"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vpc/v2/model"
	component_hc_config "go_demo/component_hc_config"
	"testing"
)

func TestQueryVpc(t *testing.T) {
	akSk := component_hc_config.ReadHcConfig()
	// 初始化客户端认证信息，使用当前客户端初始化方式可不填 projectId/domainId，以初始化 global.NewCredentialsBuilder() 为例
	globalAuth := basic.NewCredentialsBuilder().
		WithAk(akSk.Ak).
		WithSk(akSk.Sk).
		// 可不填 domainId
		Build()
	fmt.Println(globalAuth)
	vpcClient := vpc.NewVpcClient(vpc.VpcClientBuilder().WithRegion(region.CN_NORTH_4).WithCredential(globalAuth).WithHttpConfig(config.DefaultHttpConfig()).Build())
	request := model.ListVpcsRequest{
		Limit:               nil,
		Marker:              nil,
		Id:                  nil,
		EnterpriseProjectId: nil,
	}
	listVpcs, err := vpcClient.ListVpcs(&request)
	if err != nil {
		panic(err)
	}
	fmt.Println(listVpcs)
}
