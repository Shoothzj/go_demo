package main

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"go_demo/demo_gf/service"
)

func init() {
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.GET("/hello", service.Hello)
		group.DELETE("/resources/id1", service.Delete)
	})
}

func main() {
	g.Server().Run()
}
