package main

import (
	"fmt"
	"github.com/gogf/gf/util/gconv"
	"testing"
)

type Aux struct {
	Xx string
	Yy string
}

func TestMap2Struct(t *testing.T) {
	var dataMap = make(map[string]string)
	dataMap["Xx"] = "11"
	dataMap["Yy"] = "22"
	aux := Aux{}
	err := gconv.Struct(dataMap, &aux)
	if err != nil {
		panic(err)
	}
	fmt.Println(aux)
}
