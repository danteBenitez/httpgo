package httpgo

import (
	"fmt"
	"net"
	"os"
)

type Route struct {
	method  HttpMethod
	path    string
	handler HttpHandler
}

type Router struct {
	routes     []Route
	middleware []HttpMiddleware
}

func (r *Router) AddMiddleware(middleware HttpMiddleware) *Router {
	r.middleware = append(r.middleware, middleware)
	return r
}

func (r *Router) Get(path string, handler HttpHandler) *Router {
	r.routes = append(r.routes, Route{
		method: GET, 
		path: path, 
		handler: handler,
	})
	return r
}

func (r *Router) Post(path string, handler HttpHandler) *Router {
	r.routes = append(r.routes, Route{POST, path, handler})
	return r
}

func (r *Router) Put(path string, handler HttpHandler) *Router {
	r.routes = append(r.routes, Route{PUT, path, handler})
	return r
}

func (r *Router) Delete(path string, handler HttpHandler) *Router {
	r.routes = append(r.routes, Route{
		method: DELETE, 
		path: path, 
		handler: handler,
	})
	return r
}

func (r *Router) Dispatch(req *HttpRequest) *HttpResponse {
	res := NewHttpResponse().SetStatus(404, "Not Found")
	i := 0
	for {
		if i >= len(r.middleware) - 1 {
			break
		}
		fmt.Printf("Middleware %v\n (list: %v)", i, r.middleware)
		middleware := r.middleware[i]
		var next_middleware NextMiddleware 
		next_middleware = func() {
			i += 1
		}
		middleware(req, res, next_middleware)
		if res.closed {
			break
		}
	}

	for _, route := range r.routes {
		if req.method == route.method && req.path == route.path {
			route.handler(req, res)
			if res.closed {
				break
			}
		}
	}
	
	return res
}

func (a *Application) Dispatch(req *HttpRequest) *HttpResponse {
	res := NewHttpResponse().SetStatus(404, "Not Found")
	fmt.Printf("Dispatching request %v\n (routers: %v)\n", req, a.routers)
	for _, router := range a.routers {
		res = router.Dispatch(req)
	}
	fmt.Printf("Dispatched request %v\n", res)
	return res
}

func (a *Application) Serve(port int, callback func()) {
	ln, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not begin listening for connections: %v", err)
	}
	callback()
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Fprintf(os.Stderr, "could not accept connection: %v", err)
			continue
		}
		fmt.Printf("Connection accepted %v", conn)
		go func() {
			req, err := a.parseRequestFromReader(conn)
			if err != nil {
				fmt.Fprintf(os.Stderr, "could not parse request from connection: %v", err)
				conn.Close()
			}
			fmt.Printf("Request parsed %v\n", req)
			res := a.Dispatch(&req)
		    conn.Write([]byte(res.HttpString()))	
			conn.Close()
		}()
	}
}


