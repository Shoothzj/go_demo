package wal

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"testing"
)

const (
	DefaultSegmentSize = 128 * 1024 * 1024 // 128 MB
	pageSize           = 32 * 1024         // 32KB
	recordHeaderSize   = 7
)

func TestWalReader(t *testing.T) {
	f, err := os.Open("/Users/akka/Downloads/prom-data/wal/00000020")
	if err != nil {
		panic(err)
		return
	}
	parseWalLog(f)
}

func parseWalLog(f *os.File) {
	typeByte := make([]byte, 1)
	readLen, err := f.Read(typeByte)
	if err != nil {
		panic(err)
		return
	}
	if readLen != 1 {
		return
	}
	for {
		lenByte := make([]byte, 2)
		n, err := f.Read(lenByte)
		if err != nil {
			panic(err)
			return
		}
		if n != 2 {
			fmt.Println("n is ", n)
			return
		}
		length := binary.BigEndian.Uint16(lenByte)
		fmt.Println("type is ", typeByte)
		fmt.Println("len is ", length)
		readSize := pageSize - 3
		dataByte := make([]byte, readSize)
		readLen, err := f.Read(dataByte)
		if err != nil {
			panic(err)
			return
		}
		if readLen != int(readSize) {
			fmt.Println("read length  is ", readLen)
			return
		}
		fmt.Println("data is ", len(dataByte))
		parse(dataByte)
	}
}

func parse(dataByte []byte) {
	reader := bytes.NewReader(dataByte)
	idByte := make([]byte, 8)
	readLen, err := reader.Read(idByte)
	if err != nil {
		panic(err)
		return
	}
	if readLen != 8 {
		fmt.Println("read length  is ", readLen)
	}
	id := binary.BigEndian.Uint64(idByte)
	fmt.Println("id is ", id)

}
