package main

import "github.com/panjf2000/gnet"

type ExampleServer struct {
	*gnet.EventServer
}

func main() {
	gnet.Serve(&ExampleServer{}, "tcp://localhost:1998", gnet.WithCodec(&gnet.LineBasedFrameCodec{}))
}
