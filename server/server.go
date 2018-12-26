package main

import "net"
import "fmt"
import "euler/server/connection"
import "euler/server/httprequest"

func HandleConnection(conn net.Conn) {
	fmt.Println("Connection received!")
	input, err := connection.ReadFrom(conn)

	if err != nil {
		fmt.Printf("There was an error while reading from the client: %s.\n", err)
	} else {
		headers, _ := httprequest.From(input)
		fmt.Printf("Received from client:\n\n%s\n\n", input)
		fmt.Printf("%v", headers)
	}

	err = conn.Close()

	if err != nil {
		fmt.Printf("There was an error while closing the connection from the client: %s.\n", err)
	}
}

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Printf("There was an error while starting a server on port 8080 %s\n", err)
		return
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Printf("There was an error while opening a connection %s\n", err)
			continue
		}
		go HandleConnection(conn)
	}
}
