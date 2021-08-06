package main

import (
	"encoding/binary"
	"github.com/panjf2000/gnet"
)

type ExampleServer struct {
	*gnet.EventServer
}

func main() {
	encoderConfig := gnet.EncoderConfig{
		ByteOrder:                       binary.BigEndian,
		LengthFieldLength:               4,
		LengthAdjustment:                0,
		LengthIncludesLengthFieldLength: true,
	}
	decoderConfig := gnet.DecoderConfig{
		ByteOrder:           binary.BigEndian,
		LengthFieldOffset:   0,
		LengthFieldLength:   4,
		LengthAdjustment:    -4,
		InitialBytesToStrip: 4,
	}
	codec := gnet.NewLengthFieldBasedFrameCodec(encoderConfig, decoderConfig)
	gnet.Serve(&ExampleServer{}, "tcp://localhost:1998", gnet.WithCodec(codec))
}
