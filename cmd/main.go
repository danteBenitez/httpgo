package main

import (
	"fmt"
	"strings"

	"github.com/danteBenitez/httpgo"
)

func main() {
	app := httpgo.NewWithDefaults()
	router := app.Router()
	router.AddMiddleware(func (req *httpgo.HttpRequest, res *httpgo.HttpResponse, next httpgo.NextMiddleware) {
		fmt.Println("Middleware 1")
		next()
	})
	router.Get("/", func (req *httpgo.HttpRequest, res *httpgo.HttpResponse) {
		res.SetHeader("Content-Type", "text/html")
		msg := "<h1>Hello, World!</h1>"
		msg_len := len(msg)
		res.SetStatus(200, "OK").SetBody(strings.NewReader(msg), msg_len)
		res.Close()
	})

	app.Serve(8080, func () {
		fmt.Println("Server running on port 8080")
	})
}
