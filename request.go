package httpgo

import (
	"errors"
	"io"
	"strings"

	readerutils "github.com/danteBenitez/httpgo/utils"
)

type HttpRequest struct {
	method       HttpMethod
	path         string
	headers      Headers
	body         io.Reader
	http_version string
}

func (a *Application) parseRequestFromReader(r io.Reader) (HttpRequest, error) {
	method_str, err := readerutils.ReadUntilChar(r, ' ')
	if err != nil {
		http_error := errors.New("could not read method from request")
		return HttpRequest{}, errors.Join(http_error, err)
	}

	method_str = strings.TrimSpace(method_str)
	method, err := ParseHttpMethod(method_str)

	if err != nil {
		http_error := errors.New("could not parse method from request")
		return HttpRequest{}, errors.Join(http_error, err)
	}

	path, err := readerutils.ReadUntilChar(r, ' ')
	path = strings.TrimSpace(path)

	if err != nil {
		http_error := errors.New("could not read path from request")
		return HttpRequest{}, errors.Join(http_error, err)
	}

	http_version, err := readerutils.ReadUntilChar(r, '\n')
	http_version = strings.TrimSpace(http_version)

	if err != nil {
		http_error := errors.New("could not read http version from request")
		return HttpRequest{}, errors.Join(http_error, err)
	}

	headers, err := parseHeadersFromReader(r)

	if err != nil {
		http_error := errors.New("could not parse headers from request")
		return HttpRequest{}, errors.Join(http_error, err)
	}

	return HttpRequest {
		headers:      headers,
		method:       method,
		path:         path,
		http_version: http_version,
		body: r,
	}, nil
}

func parseHeadersFromReader(r io.Reader) (Headers, error) {
	headers := Headers{make(map[string]string)}
	for {
		line, err := readerutils.ReadUntilChar(r, '\n')
		if len(strings.TrimSpace(line)) == 0 {
			return headers, nil
		}
		if err != nil {
			return headers, err
		}
		key, value, _ := strings.Cut(line, ":")
		headers.contents[key] = strings.TrimSpace(value)
	}
}

func (req *HttpRequest) SetHeader(key string, value string) *HttpRequest {
	req.headers.contents[key] = value
	return req
}
