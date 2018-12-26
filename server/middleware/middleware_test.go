package middleware_test

import (
	"euler/server/httprequest"
	"euler/server/httpresponse"
	"euler/server/middleware"
	"testing"
)

func TestFirstOneResolves(t *testing.T) {
	sut := middleware.New(func(request httprequest.HTTPRequest, response httpresponse.HTTPResponse, out chan<- *httpresponse.HTTPResponse) {
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
	sut := middleware.New(func(request httprequest.HTTPRequest, response httpresponse.HTTPResponse, out chan<- *httpresponse.HTTPResponse) {
		out <- &response
	}).Then(func(request httprequest.HTTPRequest, response httpresponse.HTTPResponse, out chan<- *httpresponse.HTTPResponse) {
		out <- &httpresponse.HTTPResponse{
			Body: "Success",
		}
	})

	channel := make(chan *httpresponse.HTTPResponse)
	go sut.ResolveFor(httprequest.HTTPRequest{}, channel)

	response := <-channel

	if response == nil {
		t.Error("Did not resolve correctly. Expected Success, got nil.")
	} else if response.Body != "Success" {
		t.Errorf("Did not resolve correctly. Expected Success, got %s.", response.Body)
	}
}
