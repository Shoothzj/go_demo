package demo_base

import (
	"fmt"
	"os/exec"
	"testing"
)

func TestExec(t *testing.T) {
	command := exec.Command("/bin/bash", "-c", "export A=B;echo $A")
	output, _ := command.Output()
	fmt.Println(string(output))
}
