package httprequest

import "strings"
import "errors"

type HTTPRequest struct {
	Line    HTTPRequestLine
	Headers []HTTPRequestHeader
	Body    string
}

type HTTPRequestLine struct {
	Method      string
	Request     string
	HTTPVersion string
}

type HTTPRequestHeader struct {
	Name  string
	Value string
}

func From(input string) (HTTPRequest, error) {
	lines := strings.Split(input, "\r\n")
	httpRequestLine, err := RequestLineFrom(lines[0])

	if err != nil {
		return HTTPRequest{}, err
	}

	lines = lines[1:]
	var headers []HTTPRequestHeader
	var body string

	for i := range lines {
		headerLine := lines[i]

		if len(headerLine) == 0 && i+1 < len(lines) {
			body = strings.Join(lines[i+1:], "\r\n")
			break
		}

		header := strings.SplitAfterN(headerLine, ": ", 2)

		if len(header) != 2 {
			return HTTPRequest{}, errors.New("there was a header that did not adhere to HTTP specs")
		}

		headers = append(headers, HTTPRequestHeader{
			Name:  strings.TrimSuffix(header[0], ": "),
			Value: header[1],
		})
	}

	return HTTPRequest{
		Line:    httpRequestLine,
		Headers: headers,
		Body:    body,
	}, nil
}

func RequestLineFrom(request string) (HTTPRequestLine, error) {
	parts := strings.Split(request, " ")

	if len(parts) < 3 {
		// TODO Better error formatting
		return HTTPRequestLine{}, errors.New("the request line did not adhere to HTTP specs")
	}

	return HTTPRequestLine{
		Method:      parts[0],
		Request:     parts[1],
		HTTPVersion: parts[2],
	}, nil
}
