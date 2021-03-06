package middleware

import (
	"euler/server/httprequest"
	"euler/server/httpresponse"
)

type MiddlewareResolver func(request httprequest.HTTPRequest, response httpresponse.HTTPResponse, out chan<- *httpresponse.HTTPResponse)

type Middleware struct {
	resolveFunction MiddlewareResolver
	next            *Middleware
}

func (middleware *Middleware) ResolveFor(request httprequest.HTTPRequest, out chan<- *httpresponse.HTTPResponse) {
	middleware.resolveFor(request, nil, out)
}

func (middleware *Middleware) resolveFor(request httprequest.HTTPRequest, response *httpresponse.HTTPResponse, out chan<- *httpresponse.HTTPResponse) {
	internalChannel := make(chan *httpresponse.HTTPResponse)

	if response == nil {
		response = httpresponse.NoResponse()
	}

	go middleware.resolveFunction(request, *response, internalChannel)

	result := <-internalChannel

	if middleware.next != nil {
		middleware.next.resolveFor(request, result, out)
	} else {
		out <- result
	}
}

func New() *Middleware {
	return WithResolver(func(request httprequest.HTTPRequest, response httpresponse.HTTPResponse, out chan<- *httpresponse.HTTPResponse) {
		out <- &response
	})
}

func WithResolver(resolver MiddlewareResolver) *Middleware {
	return &Middleware{
		resolveFunction: resolver,
	}
}

func (first *Middleware) Then(resolver MiddlewareResolver) *Middleware {
	middleware := WithResolver(resolver)
	first.append(middleware)
	return first
}

func (middleware *Middleware) append(final *Middleware) {
	if middleware.next != nil {
		middleware.next.append(final)
	} else {
		middleware.next = final
	}
}

func (middleware *Middleware) Get(path string, resolver MiddlewareResolver) *Middleware {
	return middleware.Then(func(request httprequest.HTTPRequest, response httpresponse.HTTPResponse, out chan<- *httpresponse.HTTPResponse) {
		if path == request.Line.Request {
			resolver(request, response, out)
		} else {
			out <- &response
		}
	})
}
