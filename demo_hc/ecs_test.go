package demo_hc

import (
	"fmt"
	"testing"
)

func TestQueryEcs(t *testing.T) {
	queryEcs := QueryEcs()
	fmt.Println(queryEcs)
}

func TestDelEcs(t *testing.T) {
	DelEcs()
}
