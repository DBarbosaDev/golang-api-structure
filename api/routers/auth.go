package routers

import (
	"golang-api-structure/api/handlers"
	"golang-api-structure/api/middlewares"
	"golang-api-structure/services/routerService"
	"net/http"
)

func setAuthRouter(serveMux *http.ServeMux) {
	router := routerService.Router{ServeMux: serveMux}

	router.SetGlobalMiddlewares(middlewares.NormalizeData, middlewares.IsAuthenticated)

	router.
		SetMiddlewares(middlewares.IsAuthenticated).
		Get("/login", handlers.Login)

	router.Post("/regist", handlers.Regist)
}
