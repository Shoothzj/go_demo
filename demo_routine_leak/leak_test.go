package demo_routine_leak

import (
	"go.uber.org/goleak"
	"testing"
	"time"
)

func TestLeak(t *testing.T) {
	defer goleak.VerifyNone(t)
	if err := leak(); err != nil {
		t.Fatal("error not expected")
	}
	time.Sleep(time.Minute * 2)
}
