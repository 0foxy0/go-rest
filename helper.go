package gorest

func findRouteByPathAndMethod(routes []Route, path string, method string) int {
	for index, route := range routes {
		//lastRoutePathCharIndex
		lRPCI := len(route.path) - 1
		//lastPathCharIndex
		lPCI := len(path) - 1

		if route.path[lRPCI] == '/' && path[lPCI] != '/' {
			route.path = route.path[:lRPCI]
		}
		if route.path != path || method != route.method {
			continue
		}
		return index
	}
	return -1
}
