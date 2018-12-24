package httprequest

import "strings"
import "errors"

type HTTPRequest struct {
	line    HTTPRequestLine
	headers []HTTPRequestHeader
	body    string
}

type HTTPRequestLine struct {
	method      string
	request     string
	httpVersion string
}

type HTTPRequestHeader struct {
	name  string
	value string
}

func From(input string) (HTTPRequest, error) {
	lines := strings.Split(input, "\r\n")
	httpRequestLine, err := RequestLineFrom(lines[0])

	if err != nil {
		return HTTPRequest{}, err
	}

	lines = lines[1:]
	var headers []HTTPRequestHeader

	for i := range lines {
		header := strings.SplitAfterN(lines[i], ": ", 2)

		if len(header) != 2 {
			return HTTPRequest{}, errors.New("There was a header that did not adhere to HTTP specs.")
		}

		headers = append(headers, HTTPRequestHeader{
			name:  strings.TrimSuffix(header[0], ": "),
			value: header[1],
		})
	}

	return HTTPRequest{
		line:    httpRequestLine,
		headers: headers,
		body:    "",
	}, nil
}

func RequestLineFrom(request string) (HTTPRequestLine, error) {
	parts := strings.Split(request, " ")

	if len(parts) < 3 {
		// TODO Better error formatting
		return HTTPRequestLine{}, errors.New("The request line did not adhere to HTTP specs.")
	}

	return HTTPRequestLine{
		method:      parts[0],
		request:     parts[1],
		httpVersion: parts[2],
	}, nil
}
