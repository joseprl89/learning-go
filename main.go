package main

import "net"
import "fmt"
import "euler/server"

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
		go server.HandleConnection(conn)
	}
}
