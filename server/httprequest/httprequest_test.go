package httprequest

import "testing"

func TestRequestLineFromDoesNotFail(t *testing.T) {
	_, err := RequestLineFrom("GET / HTTP/1.1")
	if err != nil {
		t.Errorf("There was an error while obtaining the request line %s.", err)
	}
}

func TestRequestLineFromFailsWhenLessThanThreeSpaces(t *testing.T) {
	_, err := RequestLineFrom("GET /")
	if err == nil {
		t.Errorf("There was no error when not enough values passed to RequestLine: %s.", err)
	}
}

func TestRequestLineReadsCorrectlyValues(t *testing.T) {
	requestLine, _ := RequestLineFrom("GET / HTTP/1.1")
	if requestLine.method != "GET" || requestLine.request != "/" || requestLine.httpVersion != "HTTP/1.1" {
		t.Errorf("Request line not read correctly: %s.", requestLine)
	}
}

func TestFromFailsIfInvalidRequestLine(t *testing.T) {
	_, err := From("GET HTTP/1.1\r\nHeaderOne: Some Value\r\nHeaderTwo: Some OtherValue\r\n\r\nBody")
	if err == nil {
		t.Errorf("There was no error when not enough values passed to RequestLine: %s.", err)
	}
}

func TestFromFailsOnInvalidHeaders(t *testing.T) {
	_, err := From("GET / HTTP/1.1\r\nHeaderOne Invalid Since it has no colon")
	if err == nil {
		t.Errorf("There was no error when an invalid header was passed in: %s.", err)
	}
}

func TestFromParsesHeaders(t *testing.T) {
	request, _ := From("GET / HTTP/1.1\r\nHeaderOne: Some Value")

	if request.headers[0].name != "HeaderOne" {
		t.Errorf("The name of the second header wasnt correct, got %s, expected %s", request.headers[0].name, "HeaderOne")
	}
	if request.headers[0].value != "Some Value" {
		t.Errorf("The value of the second header wasnt correct, got %s, expected %s", request.headers[0].value, "Some Value")
	}
}

func TestFromParsesMultipleHeaders(t *testing.T) {
	request, _ := From("GET / HTTP/1.1\r\nHeaderOne: Some Value\r\nHeaderTwo: Some Other Value")

	if request.headers[0].name != "HeaderOne" {
		t.Errorf("The name of the second header wasnt correct, got %s, expected %s", request.headers[0].name, "HeaderOne")
	}
	if request.headers[0].value != "Some Value" {
		t.Errorf("The value of the second header wasnt correct, got %s, expected %s", request.headers[0].value, "Some Value")
	}

	if request.headers[1].name != "HeaderTwo" {
		t.Errorf("The name of the second header wasnt correct, got %s, expected %s", request.headers[1].name, "HeaderTwo")
	}
	if request.headers[1].value != "Some Other Value" {
		t.Errorf("The value of the second header wasnt correct, got %s, expected %s", request.headers[1].value, "Some Other Value")
	}
}

func TestFromParsesHeadersWithColonsInValue(t *testing.T) {
	request, _ := From("GET / HTTP/1.1\r\nHeaderOne: : Something")

	if request.headers[0].name != "HeaderOne" {
		t.Errorf("The name of the second header wasnt correct, got %s, expected %s", request.headers[0].name, "HeaderOne")
	}
	if request.headers[0].value != ": Something" {
		t.Errorf("The value of the second header wasnt correct, got %s, expected %s", request.headers[0].value, ": Something")
	}
}

func TestFromParsesBody(t *testing.T) {
	request, _ := From("GET / HTTP/1.1\r\nHeaderOne: : Something\r\n\r\nBody")

	if request.body != "Body" {
		t.Errorf("Expected body to be \"Body\", got \"%s\" instead.", request.body)
	}
}
