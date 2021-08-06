package channel_test

import (
	"fmt"
	"testing"
	"time"
)

func TestChannelTestCase1(t *testing.T) {
	times := time.After(5 * time.Second)
	ch := make(chan int)
	go func() {
		ch <- 1
		time.Sleep(1 * time.Second)
		ch <- 2
	}()
	for {
		select {
		case val := <-ch:
			fmt.Println(val)
		case <-times:
			fmt.Println("timeout")
			break
		}
	}
	fmt.Println("END")
}

func TestChannelTestCase2(t *testing.T) {
	times := time.After(5 * time.Second)
	ch := make(chan int)
	go func() {
		ch <- 1
		time.Sleep(1 * time.Second)
		ch <- 2
	}()
FOR:
	for {
		select {
		case val := <-ch:
			fmt.Println(val)
		case <-times:
			fmt.Println("timeout")
			break FOR
		}
	}
	fmt.Println("END")
}
