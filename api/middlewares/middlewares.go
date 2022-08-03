package middlewares

import (
	"golang-api-structure/services/routerService"
	"log"
	"net/http"
)

func IsAuthenticated(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return routerService.NewMiddleware(handlerFunc, func() bool {
		log.Println("Checking if isAuthenticated")
		return true
	}, func() {
		log.Println("IsAuthenticated finishing")
	})
}

func NormalizeData(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return routerService.NewMiddleware(handlerFunc, func() bool {
		log.Println("Normalizing body")
		return true
	}, func() {
		log.Println("NormalizeData finishing")
	})
}
