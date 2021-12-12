module go_demo/demo_hc

go 1.17

require (
	github.com/huaweicloud/huaweicloud-sdk-go-obs v3.21.1+incompatible // indirect
	github.com/huaweicloud/huaweicloud-sdk-go-v3 v0.0.39-rc
	go_demo/component_hc_config v0.0.0
)

replace go_demo/component_hc_config v0.0.0 => ../component_hc_config
