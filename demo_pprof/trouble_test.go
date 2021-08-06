package main

import (
	"testing"
	"time"
)

func TestCPU(t *testing.T) {
	for i := 0; i < 8; i++ {
		go test()
	}
	time.Sleep(time.Hour * 1)
}

func test() {
	time.AfterFunc(time.Duration(time.Second*1), func() {
		test()
	})
}
