package demo_base

import (
	"fmt"
	"testing"
)

func TestChannelSendCase1(t *testing.T) {
	ch := make(chan interface{})
	select {
	case ch <- struct{}{}:
		fmt.Println("xxxxx")
	default:
		fmt.Println("print channel")
	}
}
