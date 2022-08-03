package routers

import (
	"net/http"
)

func InitRouter(serveMux *http.ServeMux) {
	setAuthRouter(serveMux)
}
