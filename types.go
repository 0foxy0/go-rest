package gorest

import (
	"net/http"
)

type Response struct {
	http.ResponseWriter
}
type Request struct {
	*http.Request
}
type NextFunction func()
type HandlerFunc func(Response, Request, NextFunction)
type Router struct {
	routes map[string][]HandlerFunc
}
