package main

import "fmt"

func main() {
	ch := make(chan struct{})
	go func() {
		_ = <-ch
	}()
	ch <- struct{}{}
	fmt.Println("hello")
}
