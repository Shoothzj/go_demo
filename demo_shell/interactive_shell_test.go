package demo_shell

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"testing"
)

func TestCallInteractiveShell(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	cmd := exec.Command("/bin/bash", dir+"/interactive_shell.sh")
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	buffer := bytes.Buffer{}
	buffer.Write([]byte("ZhangJian"))
	cmd.Stdin = &buffer
	err = cmd.Run()
	if err != nil {
		panic(err)
	}
	outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
	fmt.Println("output is ", outStr)
	fmt.Println("err output is ", errStr)
	fmt.Println("Execute Over")
}
