package httpresponse

import (
	"euler/mocks"
	"fmt"
	"testing"
)

func TestRequestLineFromDoesNotFail(t *testing.T) {
	response := HTTPResponse{
		HttpVersion:  "HTTP/1.1",
		Status:       200,
		ReasonPhrase: "OK",
		Body:         "<html><body>Hi!</body></html>",
		Headers:      make([]Header, 0),
	}

	if response.HttpVersion != "HTTP/1.1" {
		t.Errorf("Incorrect HTTP Version %s.", response.HttpVersion)
	}

	if response.Status != 200 {
		t.Errorf("Incorrect Status code %d.", response.Status)
	}

	if response.ReasonPhrase != "OK" {
		t.Errorf("Incorrect reason phrase %s, expected OK.", response.ReasonPhrase)
	}

	if response.Body != "<html><body>Hi!</body></html>" {
		t.Errorf("Incorrect body %s, expected \"<html><body>Hi!</body></html>\".", response.ReasonPhrase)
	}

	if len(response.Headers) != 0 {
		t.Errorf("Incorrect headers %s, expected none.", response.Headers)
	}
}

func TestWriteToConnection(t *testing.T) {
	response := HTTPResponse{
		HttpVersion:  "HTTP/1.1",
		Status:       200,
		ReasonPhrase: "OK",
		Body:         "<html><body>Hi!</body></html>",
		Headers:      make([]Header, 0),
	}

	var mockedConnection = mocks.Connection{}

	err := response.WriteTo(&mockedConnection)

	writtenString := string(mockedConnection.WrittenBytes)

	expected := "HTTP/1.1 200 OK\n\n<html><body>Hi!</body></html>"

	if writtenString != expected {
		t.Errorf("Didn't write correct data to connection. Expected:\n%v\n\ngot:\n%s", expected, writtenString)
	}

	if err != nil {
		t.Error("Found an unexpected error")
	}
}

func TestWriteToConnectionWithHeaders(t *testing.T) {
	response := HTTPResponse{
		HttpVersion:  "HTTP/1.1",
		Status:       200,
		ReasonPhrase: "OK",
		Body:         "<html><body>Hi!</body></html>",
		Headers: []Header{
			{Name: "Server", Value: "LearningGoServer"},
			{Name: "Date", Value: "Today"},
		},
	}

	var mockedConnection = mocks.Connection{}

	err := response.WriteTo(&mockedConnection)

	writtenString := string(mockedConnection.WrittenBytes)

	expected := fmt.Sprintf(
		"%s\n%s\n%s\n\n%s",
		"HTTP/1.1 200 OK",
		"Server: LearningGoServer",
		"Date: Today",
		"<html><body>Hi!</body></html>")

	if writtenString != expected {
		t.Errorf("Didn't write correct data to connection. Expected:\n%v\n\ngot:\n%s", expected, writtenString)
	}

	if err != nil {
		t.Error("Found an unexpected error")
	}
}
