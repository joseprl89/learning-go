package httpresponse

import (
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"
)

type HTTPResponse struct {
	HttpVersion  string
	Status       int
	ReasonPhrase string
	Body         string
	Headers      []Header
}

type Header struct {
	Name  string
	Value string
}

func (response HTTPResponse) WriteTo(conn net.Conn) error {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("%s %d %s\n", response.HttpVersion, response.Status, response.ReasonPhrase))

	for i := range response.Headers {
		header := response.Headers[i]
		sb.WriteString(fmt.Sprintf("%s: %s\n", header.Name, header.Value))
	}

	sb.WriteString("\n")
	sb.WriteString(response.Body)

	responseBytes := []byte(sb.String())

	_, err := conn.Write(responseBytes)
	return err
}

func NoResponse() *HTTPResponse {
	return &HTTPResponse{
		Status:       500,
		ReasonPhrase: "Not implemented",
		Headers: []Header{
			{
				Name:  "Server",
				Value: "LearningGoServer",
			},
			{
				Name:  "Date",
				Value: currentTimeInRFC850Format(),
			},
		},
	}
}

func currentTimeInRFC850Format() string {
	now := time.Now()
	return now.Format(http.TimeFormat)
}
