package common

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func CheckReadLen(read, desired int) {
	if read != desired {
		panic(fmt.Sprintf("read Len is %d, desired len is %d", read, desired))
	}
}

func ReadMagic(reader *bytes.Reader) {
	magicByte := make([]byte, 4)
	readLen, err := reader.Read(magicByte)
	if err != nil {
		panic(err)
		return
	}
	fmt.Println(string(magicByte))
	CheckReadLen(readLen, 4)
}

func ReadChecksum(reader *bytes.Reader) {
	magicByte := make([]byte, 4)
	readLen, err := reader.Read(magicByte)
	if err != nil {
		panic(err)
		return
	}
	fmt.Println(string(magicByte))
	CheckReadLen(readLen, 4)
}

func ReadLen(f *bytes.Reader) uint32 {
	lengthByte := make([]byte, 4)
	readLen, err := f.Read(lengthByte)
	if err != nil {
		panic(err)
		return 0
	}
	CheckReadLen(readLen, 4)
	return binary.BigEndian.Uint32(lengthByte)
}
