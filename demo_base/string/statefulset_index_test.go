package string

import (
	"fmt"
	"strings"
	"testing"
)

func TestStatefulSet(t *testing.T) {
	str := "zookeeper-0"
	index := strings.LastIndex(str, "-")
	fmt.Println(str[0:index])
	fmt.Println(str[index:len(str)])
}
