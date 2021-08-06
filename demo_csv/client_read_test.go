package main

import (
	"fmt"
	"github.com/gocarina/gocsv"
	"os"
	"testing"
)

// Client use "-" to ignore a field
type Client struct {
	Id      string `csv:"client_id"`
	Name    string `csv:"client_name"`
	Age     string `csv:"client_age"`
	NotUsed string `csv:"-"`
}

func TestClientRead(t *testing.T) {
	clientsFile, err := os.OpenFile("client.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer clientsFile.Close()

	clients := []*Client{}

	// Load clients from file
	if err := gocsv.UnmarshalFile(clientsFile, &clients); err != nil {
		panic(err)
	}

	for _, client := range clients {
		fmt.Println("Hello", client.Name)
	}

	if _, err := clientsFile.Seek(0, 0); err != nil { // Go to the start of the file
		panic(err)
	}

	clients = append(clients, &Client{Id: "12", Name: "John", Age: "21"}) // Add clients
	clients = append(clients, &Client{Id: "13", Name: "Fred"})
	clients = append(clients, &Client{Id: "14", Name: "James", Age: "32"})
	clients = append(clients, &Client{Id: "15", Name: "Danny"})

	// Get all clients as CSV string
	csvContent, err := gocsv.MarshalString(&clients)
	if err != nil {
		panic(err)
	}
	// Display all clients as CSV string
	fmt.Println(csvContent)
}
