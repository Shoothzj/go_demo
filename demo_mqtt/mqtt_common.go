package demo_mqtt

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"time"
)

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v", err)
}

func publishMessage(client mqtt.Client, deviceId string) {

	text := `{
    "name" : "zhangjian"
}`
	token := client.Publish(fmt.Sprintf(messageUpTopic, deviceId), 0, false, text)
	token.Wait()
	time.Sleep(time.Second)
}

func publishProperty(client mqtt.Client, deviceId string) {

	text := `{
    "services": [{
            "service_id": "Tempera1ture",
            "properties": {
                "value": 50
            },
            "event_time": "20210514T010000Z"
        },
        {
            "service_id": "Battery",
            "properties": {
                "level": 40
            },
            "event_time": "20210514T010000Z"
        }
    ]
}`
	token := client.Publish(fmt.Sprintf(propertyReportTopic, deviceId), 0, false, text)
	token.Wait()
	time.Sleep(time.Second)
}

func sub(client mqtt.Client) {
	topic := "topic/test"
	token := client.Subscribe(topic, 1, nil)
	token.Wait()
	fmt.Printf("Subscribed to topic: %s", topic)
}
