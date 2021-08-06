package main

import (
	"fmt"
	"go/scanner"
	"go/token"
	"io/ioutil"
)

func main() {
	// src is the input that we want to tokenize.
	src, _ := ioutil.ReadFile(`main.go`)

	// Initialize the scanner
	var s scanner.Scanner
	// positions are relative to fSet
	fSet := token.NewFileSet()
	file := fSet.AddFile("", fSet.Base(), len(src))
	// nil means no error handler
	s.Init(file, src, nil, scanner.ScanComments)

	// Repeated calls to Scan yield the token sequence found in the input
	for {
		pos, tok, lit := s.Scan()
		if tok == token.EOF {
			break
		}
		fmt.Printf("%s\t%s\t%q\n", fSet.Position(pos), tok, lit)
	}
}
