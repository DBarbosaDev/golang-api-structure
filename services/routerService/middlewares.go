package routerService

import (
	"net/http"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

type MiddlewareMethods interface {
	SetGlobalMiddlewares([]Middleware)
	SetMiddlewares([]Middleware) *Router
	RemoveMiddlewares() *Router
}

func NewMiddleware(routeHandler http.HandlerFunc, middlewareCallback func() bool, onFinish func()) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// Implement some work BEFORE executing the route handler
		if !middlewareCallback() {
			return
		}

		// Execute the route handler
		routeHandler.ServeHTTP(writer, request)

		// Implement some work AFTER executing the route handler
		onFinish()
	}
}

func (router *Router) SetGlobalMiddlewares(middlewares ...Middleware) {
	router.GlobalMiddlewares = middlewares
}

func (router *Router) SetMiddlewares(middlewares ...Middleware) *Router {
	router.LocalMiddlewares = middlewares
	return router
}

func (router *Router) RemoveMiddlewares() *Router {
	router.LocalMiddlewares = nil
	return router
}
