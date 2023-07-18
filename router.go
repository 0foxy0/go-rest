package gorest

import (
	"fmt"
	"net/http"
)

func (r *Router) Use(path string, router *Router) {
	for key, value := range router.routes {
		route := fmt.Sprintf("%s%s", path, key)
		r.routes[route] = value
	}
}
func (r *Router) Get(path string, handlers ...HandlerFunc) {
	r.routes[path] = handlers
}

func (r *Router) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	path := req.URL.Path

	if handlers, ok := r.routes[path]; ok {
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
