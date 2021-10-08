package main

import "math/rand"

func main() {
	num := getRandom()
	println(*num)
}

//go:noinline
func getRandom() *int {
	tmp := rand.Intn(100)
	return &tmp
}
