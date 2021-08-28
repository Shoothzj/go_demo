package main

import "fmt"

type Xx struct {
	yy string
}


func main() {
	var xx Xx
	xx.yy = "yy"
	fmt.Println(xx)
	fmt.Println(xx.yy)
}
