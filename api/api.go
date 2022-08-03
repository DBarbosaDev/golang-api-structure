package api

import (
	"golang-api-structure/api/routers"
	"net/http"
)

func InitApi(serveMux *http.ServeMux) {
	routers.InitRouter(serveMux)
}
