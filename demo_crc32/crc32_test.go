package demo_crc32

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"hash/crc32"
	"testing"
)

var table = crc32.MakeTable(0x82f63b78)

func Test_Crc32(t *testing.T) {
	bytes, _ := hex.DecodeString("00000000" + "00000000" + "017a92e383dd0000017a92e383ddffffffffffffffffffffffffffff000000011c000000011053686f6f74487a6a00")
	checksumIEEE := crc32.Checksum(bytes, table)
	fmt.Println(checksumIEEE)
	auxBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(auxBytes, checksumIEEE)
	fmt.Println(auxBytes)
	encodeToString := hex.EncodeToString(auxBytes)
	fmt.Println(encodeToString)
}
