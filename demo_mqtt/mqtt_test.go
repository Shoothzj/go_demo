package demo_mqtt

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	component_hc_config "go_demo/component_hc_config"
	"testing"
)

func TestMqtt(t *testing.T) {
	hcConfig := component_hc_config.ReadHcConfig()
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", mqttDomain, 1883))
	opts.SetClientID(clientId(hcConfig.DeviceId))
	opts.SetUsername(hcConfig.DeviceId)
	opts.SetPassword(computeHmac256(hcConfig.DeviceSk, "2021051409"))
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	sub(client)
	publishProperty(client, hcConfig.DeviceId)

	client.Disconnect(250)
}
