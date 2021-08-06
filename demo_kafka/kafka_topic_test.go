package demo_kafka

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"testing"
)

func TestCreateTopic(t *testing.T) {
	// to create topics when auto.create.topics.enable='true'
	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", "my-topic", 0)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(conn)
}
