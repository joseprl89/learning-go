package middleware_test

import (
	"euler/server/httprequest"
	"euler/server/httpresponse"
	"euler/server/middleware"
	"testing"
)

func TestFirstOneResolves(t *testing.T) {
	sut := middleware.WithResolver(func(request httprequest.HTTPRequest, response httpresponse.HTTPResponse, out chan<- *httpresponse.HTTPResponse) {
		out <- &httpresponse.HTTPResponse{
			Body: "Success",
		}
	})

	channel := make(chan *httpresponse.HTTPResponse)
	go sut.ResolveFor(httprequest.HTTPRequest{}, channel)

	response := <-channel

	if response.Body != "Success" {
		t.Errorf("Did not resolve correctly. Expected Success, got %s.", response.Body)
	}
}

func TestSecondOneResolves(t *testing.T) {
	sut := middleware.WithResolver(func(request httprequest.HTTPRequest, response httpresponse.HTTPResponse, out chan<- *httpresponse.HTTPResponse) {
		out <- &response
	}).Then(func(request httprequest.HTTPRequest, response httpresponse.HTTPResponse, out chan<- *httpresponse.HTTPResponse) {
		out <- &httpresponse.HTTPResponse{
			Body: "Success",
		}
	})

	channel := make(chan *httpresponse.HTTPResponse)
	go sut.ResolveFor(httprequest.HTTPRequest{}, channel)

	response := <-channel

	if response.Body != "Success" {
		t.Errorf("Did not resolve correctly. Expected Success, got %s.", response.Body)
	}
}

func TestUntreatedReturns500(t *testing.T) {
	sut := middleware.WithResolver(func(request httprequest.HTTPRequest, response httpresponse.HTTPResponse, out chan<- *httpresponse.HTTPResponse) {
		out <- &response
	})

	channel := make(chan *httpresponse.HTTPResponse)

	go sut.ResolveFor(httprequest.HTTPRequest{}, channel)

	response := <-channel

	if response.Status != 500 {
		t.Errorf("Did not resolve correctly. Expected 500 status code, got %d.", response.Status)
	}

	if response.ReasonPhrase != "Not implemented" {
		t.Errorf("Did not resolve correctly. Expected Not implemented reason phrase, got %s.", response.ReasonPhrase)
	}
}

func TestAddsHeaders(t *testing.T) {
	sut := middleware.WithResolver(func(request httprequest.HTTPRequest, response httpresponse.HTTPResponse, out chan<- *httpresponse.HTTPResponse) {
		out <- &response
	})

	channel := make(chan *httpresponse.HTTPResponse)

	go sut.ResolveFor(httprequest.HTTPRequest{}, channel)

	response := <-channel

	if len(response.Headers) < 2 {
		t.Error("Expected headers, got none.")
	}
}

func TestGetRoute(t *testing.T) {
	request := httprequest.HTTPRequest{
		Line: httprequest.HTTPRequestLine{
			Request: "/hello",
		},
	}

	helloResolver := func(request httprequest.HTTPRequest, response httpresponse.HTTPResponse, out chan<- *httpresponse.HTTPResponse) {
		response.Body = "hello!"
		out <- &response
	}

	goodbyeResolver := func(request httprequest.HTTPRequest, response httpresponse.HTTPResponse, out chan<- *httpresponse.HTTPResponse) {
		response.Body = "goodbye!"
		out <- &response
	}

	sut := middleware.New().
		Get("/", goodbyeResolver).
		Get("/hello", helloResolver).
		Get("/goodbye", goodbyeResolver)

	channel := make(chan *httpresponse.HTTPResponse)

	go sut.ResolveFor(request, channel)

	response := <-channel

	if response.Body != "hello!" {
		t.Errorf("Expected body \"hello!\", got \"%s\".", response.Body)
	}
}
