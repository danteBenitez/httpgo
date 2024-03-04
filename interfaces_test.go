package httpgo

import "testing"

func TestParsesHttpMethodCorrectly(t *testing.T) {
	method, err := ParseHttpMethod("GET")
	if err != nil {
		t.Error("Expected no error, got ", err)
	}
	if method != GET {
		t.Error("Expected GET, got ", method)
	}

}
