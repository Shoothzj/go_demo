package main

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
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
