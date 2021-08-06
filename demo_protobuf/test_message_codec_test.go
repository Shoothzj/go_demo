package demo_protobuf

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"go_demo/demo_protobuf/test"
	"testing"
)

func Test_TestMessageCodec(t *testing.T) {
	var a int32 = 150
	testMsg := &test.Test1{A: &a}
	bytes, err := proto.Marshal(testMsg)
	if err != nil {
		panic(err)
	}
	fmt.Println(bytes)
}
