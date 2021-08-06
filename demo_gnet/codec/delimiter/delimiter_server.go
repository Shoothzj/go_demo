package main

import "github.com/panjf2000/gnet"

type ExampleServer struct {
	*gnet.EventServer
}

func main() {
	codec := gnet.NewDelimiterBasedFrameCodec(0x11)
	gnet.Serve(&ExampleServer{}, "tcp://localhost:1998", gnet.WithCodec(codec))
}
