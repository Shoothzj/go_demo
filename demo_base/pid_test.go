package demo_base

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestPid(t *testing.T) {
	fmt.Println(os.Getpid())
	go func() {
		fmt.Println(os.Getpid())
	}()
	go func() {
		fmt.Println(os.Getpid())
	}()
	time.Sleep(time.Second)
}
