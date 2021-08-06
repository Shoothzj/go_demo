package demo_base

import (
	"context"
	"fmt"
	"net"
	"testing"
	"time"
)

func TestDnsResolver(t *testing.T) {
	var r net.Resolver
	const timeout = 10 * time.Second
	ctx, cancel := context.WithTimeout(context.TODO(), timeout)
	defer cancel()
	host, err := r.LookupHost(ctx, "www.baidu.com")
	if err != nil {
		panic(err)
	}
	fmt.Println(host)
}
