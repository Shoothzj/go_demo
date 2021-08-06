package demo_base

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	var metadataMap map[string]string = make(map[string]string)
	metadataMap["1"] = "2"
	for key, value := range metadataMap {
		fmt.Println(key)
		fmt.Println(value)
	}
}
