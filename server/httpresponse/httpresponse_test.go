package httpresponse

import (
	"euler/mocks"
	"testing"
)

func TestRequestLineFromDoesNotFail(t *testing.T) {
	response := HTTPResponse{
		httpVersion:  "HTTP/1.1",
		status:       200,
		reasonPhrase: "OK",
		body:         "<html><body>Hi!</body></html>",
		headers:      make([]HTTPResponseHeader, 0),
	}

	if response.httpVersion != "HTTP/1.1" {
		t.Errorf("Incorrect HTTP Version %s.", response.httpVersion)
	}

	if response.status != 200 {
		t.Errorf("Incorrect Status code %d.", response.status)
	}

	if response.reasonPhrase != "OK" {
		t.Errorf("Incorrect reason phrase %s, expected OK.", response.reasonPhrase)
	}

	if response.body != "<html><body>Hi!</body></html>" {
		t.Errorf("Incorrect body %s, expected \"<html><body>Hi!</body></html>\".", response.reasonPhrase)
	}

	if len(response.headers) != 0 {
		t.Errorf("Incorrect headers %s, expected none.", response.headers)
	}
}

func TestWriteToConnection(t *testing.T) {
	response := HTTPResponse{
		httpVersion:  "HTTP/1.1",
		status:       200,
		reasonPhrase: "OK",
		body:         "<html><body>Hi!</body></html>",
		headers:      make([]HTTPResponseHeader, 0),
	}

	mockedConnection := mocks.Connection{}

	response.WriteTo(mockedConnection)

	if len(mockedConnection.WrittenBytes) == 0 {
		t.Errorf("Didn't write anything to connection, expected: \"%v\"", response)
	}
}
