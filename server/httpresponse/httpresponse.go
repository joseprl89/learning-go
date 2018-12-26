package httpresponse

import (
	"fmt"
	"net"
	"strings"
)

type HTTPResponse struct {
	httpVersion  string
	status       int
	reasonPhrase string
	body         string
	headers      []HTTPResponseHeader
}

type HTTPResponseHeader struct {
	name  string
	value string
}

func (response HTTPResponse) WriteTo(conn net.Conn) error {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("%s %d %s\n", response.httpVersion, response.status, response.reasonPhrase))

	for i := range response.headers {
		header := response.headers[i]
		sb.WriteString(fmt.Sprintf("%s: %s\n", header.name, header.value))
	}

	sb.WriteString("\n")
	sb.WriteString(response.body)

	responseBytes := []byte(sb.String())

	_, err := conn.Write(responseBytes)
	return err
}
