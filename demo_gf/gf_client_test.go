package main

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"testing"
)

func TestGfClient(t *testing.T) {
	response, err := g.Client().Get("http://www.baidu.com")
	if err != nil {
		panic(err)
	}
	fmt.Println(response.Status)
	fmt.Println(response.StatusCode)
	fmt.Println(response.ReadAllString())
}

func TestGfClient2(t *testing.T) {
	var (
		err error
		res *ghttp.ClientResponse
	)
	res, err = g.Client().Get("http://www.baidu.com")
	if err != nil {
		panic(err)
	}
	fmt.Println(res.Status)
	fmt.Println(res.StatusCode)
	fmt.Println(res.ReadAllString())
}


