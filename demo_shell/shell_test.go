package demo_shell

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"testing"
)

func TestCallShell(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	cmd := exec.Command("/bin/bash", dir+"/shell.sh")
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err = cmd.Run()
	if err != nil {
		panic(err)
	}
	outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
	fmt.Println("output is ", outStr)
	fmt.Println("err output is ", errStr)
	fmt.Println("Execute Over")
}
