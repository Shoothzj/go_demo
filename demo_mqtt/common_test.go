package demo_mqtt

import (
	"fmt"
	"testing"
)

func TestClientId(t *testing.T) {
	id := clientId("deviceId")
	fmt.Println(id)
}
