package server

import (
	"euler/server/httprequest"
	"euler/server/httpresponse"
	"fmt"
	"net"
)

func HandleConnection(connection net.Conn) {
	defer closeConnectionAndLog(connection)

	request, readError := readClientRequest(connection)

	if readError != nil {
		fmt.Printf("There was an error while reading from the client: %s.\n", readError)
		return
	}

	sendResponseAndLog(responseFor(request), connection)
	fmt.Println("Request served.")
}

func responseFor(request httprequest.HTTPRequest) httpresponse.HTTPResponse {
	return httpresponse.HTTPResponse{
		HttpVersion:  "HTTP/1.1",
		Status:       200,
		ReasonPhrase: "OK",
		Body:         "<html><body>Hi!</body></html>",
		Headers: []httpresponse.Header{
			{Name: "Server", Value: "LearningGoServer"},
			{Name: "Date", Value: "Today"},
		},
	}
}

func readClientRequest(connection net.Conn) (httprequest.HTTPRequest, error) {
	input, readError := readFrom(connection)

	if readError != nil {
		return httprequest.HTTPRequest{}, readError
	}

	return httprequest.From(input)
}

func closeConnectionAndLog(connection net.Conn) {
	closeError := connection.Close()
	if closeError != nil {
		fmt.Printf("There was an error while closing the connection from the client: %s.\n", closeError)
	}
}

func sendResponseAndLog(response httpresponse.HTTPResponse, connection net.Conn) {
	writeError := response.WriteTo(connection)
	if writeError != nil {
		fmt.Printf("There was an error while closing the connection from the client: %s.\n", writeError)
	}
}

func readFrom(conn net.Conn) (string, error) {
	bufferSize := 32
	buffer := make([]byte, bufferSize)
	var result []byte
	for {
		n, err := conn.Read(buffer)

		if err != nil {
			return "", err
		} else {
			result = append(result, buffer[0:n]...)
		}

		if n < bufferSize {
			return string(result[:]), nil
		}
	}
}
