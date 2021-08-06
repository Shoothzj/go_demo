package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func http_call() {
	resp, err := http.Get("http://example.com/")
	if err != nil {
		return
	}
	defer resp.Body.Close()
	exampleBytes, err := ioutil.ReadAll(resp.Body)
	exampleStr := string(exampleBytes)
	fmt.Println(exampleStr)
}
