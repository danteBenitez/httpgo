package httpgo

import (
	"fmt"
	"io"
	"strconv"
	"strings"
)

type HttpResponse struct {
	status_code  int
	status_text  string
	headers      Headers
	body         io.Reader
	http_version string

	closed bool
}

func NewHttpResponse() *HttpResponse {
	return &HttpResponse{
		headers:      Headers{contents: make(map[string]string)},
		http_version: "HTTP/1.1",
		body:         strings.NewReader(""),
	}
}

func (res *HttpResponse) SetBody(body io.Reader, len int) *HttpResponse {
	res.SetHeader("Content-Length", strconv.Itoa(len))
	res.body = body
	return res
}

func (res *HttpResponse) Close() *HttpResponse{
	res.closed = true
	return res
}

func (res *HttpResponse) HttpString() string {
	builder := strings.Builder{}
	builder.Write([]byte(fmt.Sprintf("%s %d %s\n", res.http_version, res.status_code, res.status_text)))
	for key, value := range res.headers.contents {
		builder.Write([]byte(fmt.Sprintf("%s: %s\r\n", key, value)))
	}
	content_length_str, ok := res.headers.contents["Content-Length"]
	content_length, parse_error := strconv.Atoi(content_length_str)
	if !ok || parse_error != nil {
		content_length = 0
	}
	builder.Write([]byte("\r\n"))
	buf := make([]byte, content_length)
	res.body.Read(buf)
	builder.Write(buf)
	return builder.String()
}

func (res *HttpResponse) SetStatus(code int, text string) *HttpResponse {
	res.status_code = code
	res.status_text = text
	return res
}

func (res *HttpResponse) SetHeader(key, value string) *HttpResponse {
	res.headers.contents[key] = value
	return res
}
