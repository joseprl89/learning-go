package httpresponse

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