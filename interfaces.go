package httpgo

import (
	"errors"
	"strings"
)

type HttpMethod int

const (
	INVALID_METHOD HttpMethod = iota
	GET
	POST
	PUT
	PATCH
	DELETE
	HEAD
)

// Tries to parse an HTTP method from `method`.
// Returns INVALID_METHOD as it first return value
// when encountering a method not invalid or unsupported
func ParseHttpMethod(method string) (HttpMethod, error) {
	trimmed := strings.TrimSpace(method)
	methods := map[string]HttpMethod{
		"GET":    GET,
		"POST":   POST,
		"PATCH":  PATCH,
		"DELETE": DELETE,
		"HEAD":   HEAD,
	}

	parsed, ok := methods[trimmed]
	if !ok {
		return INVALID_METHOD, errors.New("method invalid or not supported")
	}

	return parsed, nil
}

type Headers struct {
	contents map[string]string
}

type HttpHandler func(req *HttpRequest, res *HttpResponse)
type NextMiddleware func()
type HttpMiddleware func(req *HttpRequest, res *HttpResponse, next NextMiddleware)
