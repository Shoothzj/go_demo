package main

import (
	"bytes"
	"encoding/csv"
	"io"
	"log"
)

var data = []byte(
	`a, b,c ,
	foo,
	1,2,3
	,",
`)

func main() {
	reader := csv.NewReader(bytes.NewBuffer(data))
	for {
		if _, err := reader.Read(); err != nil {
			if err == io.EOF {
				break
			}
			log.Printf(err.Error())
		}
	}
}

