package main

import (
	"fmt"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", "localhost:1997")
	if err != nil {
		panic(err)
	}
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			panic(err)
		}
		go handleRequest(conn)
	}
}

// handle incoming requests
func handleRequest(conn net.Conn) {
	// make a buffer to hold incoming data
	buf := make([]byte, 1024)
	// Read the incoming connection into the buffer
	reqLen, err := conn.Read(buf)
	if err != nil {
		fmt.Println("error reading: ", err.Error())
	}
	if reqLen != 10 {
		fmt.Println("invalid request size ", reqLen)
	}
	_, err = conn.Write([]byte("That's ok"))
	if err != nil {
		fmt.Println("error sending: ", err.Error())
	}
}
