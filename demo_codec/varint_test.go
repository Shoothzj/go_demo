package demo_codec

import (
	"encoding/binary"
	"fmt"
	"testing"
)

func Test_ReadUnsignedVarInt0x10(t *testing.T) {
	byteSlice := make([]byte, 0)
	byteSlice = append(byteSlice, 0x10)
	i := 0
	var value uint32 = 0
	var aux uint32
	//for {
	aux = uint32(byteSlice[i])
	fmt.Println(aux & 0x80)
	//}
	value |= (aux << i)
	fmt.Println(value)
	ux := int((value >> 1) ^ uint32((-int(value) & 1)))
	fmt.Println(ux)
}

func Test_ReadUnsignedVarInt0x01(t *testing.T) {
	byteSlice := make([]byte, 0)
	byteSlice = append(byteSlice, 0x01)
	i := 0
	var value uint32 = 0
	var aux uint32
	//for {
	aux = uint32(byteSlice[i])
	fmt.Println(aux & 0x80)
	//}
	value |= (aux << i)
	fmt.Println(value)
	ux := int((value >> 1) ^ uint32((-int(value) & 1)))
	fmt.Println(ux)
}

func Test_ReadUnsignedVarIntWay2(t *testing.T) {
	byteSlice := make([]byte, 0)
	byteSlice = append(byteSlice, 0x01)
	varint, i := binary.Varint(byteSlice)
	fmt.Println(varint)
	fmt.Println(i)
}
