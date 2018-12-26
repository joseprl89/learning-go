package middleware

import (
	"euler/server/httprequest"
	"euler/server/httpresponse"
)

type MiddlewareResolver func(request httprequest.HTTPRequest, response httpresponse.HTTPResponse, out chan<- httpresponse.HTTPResponse)

type Middleware struct {
	ResolveFunction MiddlewareResolver
	Next            *Middleware
}

func (middleware *Middleware) ResolveFor(request httprequest.HTTPRequest, out chan<- httpresponse.HTTPResponse) {
	internalChannel := make(chan httpresponse.HTTPResponse)

	go middleware.ResolveFunction(request, httpresponse.HTTPResponse{}, internalChannel)

	result := <-internalChannel
	out <- result
}

func New(resolver MiddlewareResolver) Middleware {
	return Middleware{
		ResolveFunction: resolver,
	}
}
