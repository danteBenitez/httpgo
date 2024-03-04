package httpgo

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestPrintsResponseCorrectly(t *testing.T) {
	response := NewHttpResponse()
	msg := "<h1>Hello, World!</h1>"
	response.SetStatus(200, "OK")
	response.SetHeader("Content-Type", "text/html")
	response.SetHeader("Content-Length", fmt.Sprintf("%v", len(msg)))
	response.SetBody(strings.NewReader(msg))
	expected := fmt.Sprintf("HTTP/1.1 200 OK\nContent-Type: text/html\r\nContent-Length: %v\r\n\r\n<h1>Hello, World!</h1>", len(msg))
	actual := response.HttpString()

	if actual != expected {
		t.Errorf("Expected %s but got %s", strconv.Quote(actual), strconv.Quote(expected))
	}

}
