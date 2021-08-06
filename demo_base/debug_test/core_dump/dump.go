package main

import "math/rand"

var sum int

func main() {
	for {
		n := rand.Intn(1e6)
		sum += n
		if sum%42 == 0 {
			panic(":(")
		}
	}
}
