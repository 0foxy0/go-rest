package gorest

func CreateRouter() *Router {
	return &Router{
		routes: make(map[string][]HandlerFunc),
	}
}
