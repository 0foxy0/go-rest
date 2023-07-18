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
type Route struct {
	path     string
	method   string
	handlers []HandlerFunc
}
type Router struct {
	routes []Route
}
