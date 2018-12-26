package server

import (
	"euler/server/connection"
	"euler/server/httprequest"
	"euler/server/httpresponse"
	"fmt"
	"net"
)

func HandleConnection(conn net.Conn) {
	response := httpresponse.HTTPResponse{
		HttpVersion:  "HTTP/1.1",
		Status:       200,
		ReasonPhrase: "OK",
		Body:         "<html><body>Hi!</body></html>",
		Headers: []httpresponse.Header{
			{Name: "Server", Value: "LearningGoServer"},
			{Name: "Date", Value: "Today"},
		},
	}

	fmt.Println("Connection received!")
	input, err := connection.ReadFrom(conn)

	if err != nil {
		fmt.Printf("There was an error while reading from the client: %s.\n", err)
	} else {
		headers, _ := httprequest.From(input)
		fmt.Printf("Received from client:\n\n%s\n\n", input)
		fmt.Printf("%v", headers)
	}

	err = response.WriteTo(conn)
	if err != nil {
		fmt.Printf("There was an error while closing the connection from the client: %s.\n", err)
	}

	err = conn.Close()
	if err != nil {
		fmt.Printf("There was an error while closing the connection from the client: %s.\n", err)
	}
}
