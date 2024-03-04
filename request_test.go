package httpgo

import (
	"strings"
	"testing"
)

func TestParsesRequestFromReader(t *testing.T) {
	app := NewWithDefaults()
	request, err := app.parseRequestFromReader(strings.NewReader("GET / HTTP/1.1\nHost: localhost:3000\n\n"))
	if err != nil {
		t.Error("Expected no error, got ", err)
	}
	if request.method != GET {
		t.Error("Expected GET, got ", request.method)
	}
	if request.path != "/" {
		t.Error("Expected /, got ", request.path)
	}
	if request.http_version != "HTTP/1.1" {
		t.Error("Expected HTTP/1.1, got ", request.http_version)
	}
	if request.headers.contents["Host"] != "localhost:3000" {
		t.Error("Expected localhost:3000, got ", request.headers.contents["Host"])
	}
}
