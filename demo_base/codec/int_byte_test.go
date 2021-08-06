package codec

import (
	"encoding/binary"
	"fmt"
	"testing"
)

func Test4Bytes2Int(t *testing.T) {
	bytes := make([]byte, 4)
	bytes[0] = 0x7f
	bytes[1] = 0xff
	bytes[2] = 0xff
	bytes[3] = 0xff
	u := binary.BigEndian.Uint32(bytes)
	fmt.Println(u)
}

func TestInt24Bytes(t *testing.T) {
	val := 2147483647
	bytes := make([]byte, 4)
	binary.BigEndian.PutUint32(bytes, uint32(val))
	fmt.Println(bytes)
}
