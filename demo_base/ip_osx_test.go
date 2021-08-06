package demo_base

import (
	"fmt"
	"testing"
)

func TestOsx(t *testing.T) {
	interfaceIpv4Addr, _ := GetInterfaceIpv4Addr("lo0")
	fmt.Println(interfaceIpv4Addr)
}
