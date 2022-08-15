package router

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Route struct {
	Path string
	Method string
	Handler http.Handler
}

type Router struct {
	router *httprouter.Router
}

func New(routes []Route) *Router {
	router := httprouter.New()
	for _, route := range routes {
		router.Handler(route.Method, route.Path, route.Handler)
	}

	return &Router{
		router: router,
	}
}

func (r Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.router.ServeHTTP(w, req)
}
