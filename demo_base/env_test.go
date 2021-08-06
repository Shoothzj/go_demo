package demo_base

import (
	"fmt"
	"os"
	"testing"
)

func TestGetEnv(t *testing.T) {
	getenv := os.Getenv("name")
	fmt.Println(getenv)
}
