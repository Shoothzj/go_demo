package main

import "fmt"

var AppVersion string
var AppName = "app"

func main() {
	fmt.Println(`appName: ` + AppName)
	fmt.Println(`Version: ` + AppVersion)
}
