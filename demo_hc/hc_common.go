package demo_hc

import (
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
	"go_demo/component_hc_config"
)

func getAuth() *basic.Credentials {
	akSk := component_hc_config.ReadHcConfig()
	// 初始化客户端认证信息，使用当前客户端初始化方式可不填 projectId/domainId，以初始化 global.NewCredentialsBuilder() 为例
	auth := basic.NewCredentialsBuilder().
		WithAk(akSk.Ak).
		WithSk(akSk.Sk).
		// 可不填 domainId
		Build()
	fmt.Println(auth)
	return &auth
}
