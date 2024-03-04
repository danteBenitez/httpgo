package httpgo

import (
	"time"
)

type ApplicationConfig struct {
	timeout  time.Duration
	buf_size uint
}

type Application struct {
	routers []*Router
	config  ApplicationConfig
}

func defaultConfig() ApplicationConfig {
	return ApplicationConfig{
		timeout: 0,
	}
}

func NewWithDefaults() Application {
	return Application{routers: make([]*Router, 0), config: defaultConfig()}
}

// Constructs a router for the current application
// and inmmediately appends it to the list of routers
// and returns it
func (a *Application) Router() *Router {
	routes := make([]Route, 0)
	router := Router{routes, make([]HttpMiddleware, 0)}
	a.routers = append(a.routers, &router)
	return &router
}
