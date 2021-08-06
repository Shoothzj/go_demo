package demo_base

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"testing"
)

func TestFileWalk(t *testing.T) {
	file := "/Users/akka/go/src/go_demo"
	filepath.Walk(file, func(path string, info fs.FileInfo, err error) error {
		if !info.IsDir() {
			fmt.Println(path[len(file)+1:])
		}
		return nil
	})
}
