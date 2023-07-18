package gorest

import (
	"fmt"
	"net/http"
)

func (r *Router) Use(path string, router *Router) {
	for _, route := range router.routes {
		route.path = fmt.Sprintf("%s%s", path, route.path)
		r.routes = append(r.routes, route)
	}
}
func (r *Router) Get(path string, handlers ...HandlerFunc) {
	r.routes = append(r.routes, Route{path: path, method: "GET", handlers: handlers})
}

func (r *Router) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	routeIndex := findRouteByPathAndMethod(r.routes, path, req.Method)

	if routeIndex != -1 {
		handlers := r.routes[routeIndex].handlers
		handlerIndex := 0

		var next func()
		next = func() {
			if handlerIndex >= len(handlers) {
				return
			}
			handlerIndex++
			handlers[handlerIndex](Response{rw}, Request{req}, next)
		}

		handlers[0](Response{rw}, Request{req}, next)
	} else {
		http.NotFound(rw, req)
	}
}

func (r *Router) Listen(port int, callback func()) {
	go callback()
	http.ListenAndServe(fmt.Sprintf(":%d", port), r)
}
