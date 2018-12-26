package middleware

import (
	"euler/server/httprequest"
	"euler/server/httpresponse"
)

type MiddlewareResolver func(request httprequest.HTTPRequest, response httpresponse.HTTPResponse, out chan<- httpresponse.HTTPResponse)

type Middleware struct {
	resolveFunction MiddlewareResolver
	next            *Middleware
}

func (middleware *Middleware) ResolveFor(request httprequest.HTTPRequest, out chan<- httpresponse.HTTPResponse) {
	internalChannel := make(chan httpresponse.HTTPResponse)

	go middleware.resolveFunction(request, httpresponse.HTTPResponse{}, internalChannel)

	result := <-internalChannel
	out <- result
}

func New(resolver MiddlewareResolver) Middleware {
	return Middleware{
		resolveFunction: resolver,
	}
}
