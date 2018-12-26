package middleware

import (
	"euler/server/httprequest"
	"euler/server/httpresponse"
)

type Middleware struct {
	ResolveFunction func(request httprequest.HTTPRequest, response httpresponse.HTTPResponse, out chan<- httpresponse.HTTPResponse)
	Next            *Middleware
}

func (middleware *Middleware) ResolveFor(request httprequest.HTTPRequest, out chan<- httpresponse.HTTPResponse) {
	internalChannel := make(chan httpresponse.HTTPResponse)

	go middleware.ResolveFunction(request, httpresponse.HTTPResponse{}, internalChannel)

	result := <-internalChannel
	out <- result
}
