package routerService

import (
	"io"
	"net/http"
)

type Router struct {
	ServeMux          *http.ServeMux
	GlobalMiddlewares []Middleware
	LocalMiddlewares  []Middleware
}

type RequestMethods interface {
	Get(string, http.HandlerFunc)
	Post(string, http.HandlerFunc)
	Put(string, http.HandlerFunc)
	Delete(string, http.HandlerFunc)
}

func routeHandler(method string, handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != method {
			io.WriteString(writer, "404 not found")
			return
		}

		if handlerFunc == nil {
			return
		}

		handlerFunc(writer, request)
	}
}

func setRoute(router *Router, path string, method string, handlerFunc http.HandlerFunc) {
	middlewares := append(router.GlobalMiddlewares, router.LocalMiddlewares...)

	if len(middlewares) < 1 {
		router.ServeMux.HandleFunc(path, routeHandler(method, handlerFunc))
		return
	}

	wrapped := handlerFunc

	// loop in reverse to preserve middleware order
	for i := len(middlewares) - 1; i >= 0; i-- {
		wrapped = middlewares[i](wrapped)
	}

	router.ServeMux.HandleFunc(path, routeHandler(method, wrapped))

	router.RemoveMiddlewares()
}

func (router *Router) Get(path string, handler http.HandlerFunc) {
	setRoute(router, path, http.MethodGet, handler)
}

func (router *Router) Post(path string, handler http.HandlerFunc) {
	setRoute(router, path, http.MethodPost, handler)
}

func (router *Router) Put(path string, handler http.HandlerFunc) {
	setRoute(router, path, http.MethodPut, handler)
}

func (router *Router) Delete(path string, handler http.HandlerFunc) {
	setRoute(router, path, http.MethodDelete, handler)
}
