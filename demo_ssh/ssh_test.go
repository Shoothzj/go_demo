package demo_ssh

import (
	"fmt"
	"github.com/melbahja/goph"
	"testing"
)

func TestSsh(t *testing.T) {
	auth, err := goph.Key("", "")
	if err != nil {
		panic(err)
	}
	client, err := goph.New("root", "127.0.0.1", auth)
	if err != nil {
		panic(err)
	}
	// defer closing the network connection
	defer client.Close()

	out, err := client.Run("ls /tmp/")

	if err != nil {
		panic(err)
	}

	fmt.Println(string(out))
}
