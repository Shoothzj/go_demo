package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"testing"
)

func TestReadIris(t *testing.T) {
	fmt.Println(os.Getwd())
	// Open the CSV
	f, err := os.Open("iris.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// create a new csv reader reading from the opened file
	reader := csv.NewReader(f)

	// Assume we don't know the number of fields per line. By setting
	// FieldsPerRecord negative, each row may have a variable number of fields
	reader.FieldsPerRecord = -1
	// Read in all of the CSV Records
	rawCSVData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(rawCSVData)
}
