package dirtest

import (
	"fmt"
	"log"
	"os/user"
	"testing"
)

func TestHomeDir(t *testing.T) {
	user, err := user.Current()
	if err != nil {
		log.Fatalf(err.Error())
	}
	homeDirectory := user.HomeDir
	fmt.Printf("Home Directory: %s\n", homeDirectory)
}
