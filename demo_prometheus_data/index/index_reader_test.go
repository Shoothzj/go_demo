package index

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"go_demo/demo_prometheus-data/common"
	"io/ioutil"
	"testing"
)

func TestWalReader(t *testing.T) {
	data, err := ioutil.ReadFile("/Users/akka/Downloads/prom-data/01F0FCRBC3N5QJQT1PHTG0QJ9P/index")
	if err != nil {
		panic(err)
		return
	}
	parseIndexLog(bytes.NewReader(data))
}

func parseIndexLog(reader *bytes.Reader) {
	var offset int64
	var readLen int
	var err error
	common.ReadMagic(reader)
	offset += 4
	versionByte := make([]byte, 1)
	readLen, err = reader.Read(versionByte)
	if err != nil {
		panic(err)
		return
	}
	common.CheckReadLen(readLen, 1)
	offset += 1
	fmt.Println("version is ", versionByte)

	symbolLength := common.ReadLen(reader)
	fmt.Println(symbolLength)

	symbolLength = common.ReadLen(reader)
	fmt.Println(symbolLength)

	for i := 0; i < int(symbolLength); i++ {
		readUvarint, err := binary.ReadUvarint(reader)
		if err != nil {
			panic(err)
			return
		}
		stringByte := make([]byte, readUvarint)
		readLen, err = reader.Read(stringByte)
		if err != nil {
			panic(err)
			return
		}
		common.CheckReadLen(readLen, int(readUvarint))
		fmt.Printf("len is %d val is %s\n", readUvarint, string(stringByte))
	}
	common.ReadChecksum(reader)

}
