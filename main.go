package gorest

func CreateRouter() *Router {
	return &Router{
		routes: make([]Route, 0),
	}
}
