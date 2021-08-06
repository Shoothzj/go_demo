package demo_pulsar

import (
	"fmt"
	"github.com/apache/pulsar-client-go/pulsar"
	"log"
	"reflect"
	"testing"
)

// TestMultiConsumer
func TestMultiConsumer(t *testing.T) {
	client, err := pulsar.NewClient(pulsar.ClientOptions{URL: "pulsar://localhost:6650"})
	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()

	consumer1, err := client.Subscribe(pulsar.ConsumerOptions{
		Topic:            "topic-1",
		SubscriptionName: "sub",
	})
	consumer2, err := client.Subscribe(pulsar.ConsumerOptions{
		Topic:            "topic-2",
		SubscriptionName: "sub",
	})
	if err != nil {
		log.Fatal(err)
	}
	defer consumer1.Close()
	defer consumer2.Close()

	chanSlice := make([]<-chan pulsar.ConsumerMessage, 0)
	chanSlice = append(chanSlice, consumer1.Chan())
	chanSlice = append(chanSlice, consumer2.Chan())

	reflectSelectSlice := make([]reflect.SelectCase, 0)
	reflectSelectSlice = append(reflectSelectSlice, reflect.SelectCase{
		Dir:  reflect.SelectRecv,
		Chan: reflect.ValueOf(consumer1.Chan()),
	})
	reflectSelectSlice = append(reflectSelectSlice, reflect.SelectCase{
		Dir:  reflect.SelectRecv,
		Chan: reflect.ValueOf(consumer2.Chan()),
	})
	for {
		chosen, recv, _ := reflect.Select(reflectSelectSlice)
		ch := chanSlice[chosen]
		fmt.Println("choosen chan is ", ch)
		consumerMessage := recv.Interface().(pulsar.ConsumerMessage)
		fmt.Println("consumer message is ", consumerMessage.Message.PublishTime())
	}
}
