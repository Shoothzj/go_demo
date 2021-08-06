module go_demo/demo_mqtt

go 1.16

require (
	github.com/eclipse/paho.mqtt.golang v1.3.4
	go_demo/component_hc_config v0.0.0
)

replace go_demo/component_hc_config v0.0.0 => ../component_hc_config
