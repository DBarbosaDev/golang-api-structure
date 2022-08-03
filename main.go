package main

import (
	"golang-api-structure/api"
	"golang-api-structure/configs"
	"log"
	"net/http"
)

func main() {
	envConfigs := configs.GetEnvConfigs()
	serverMux := http.NewServeMux() // or &http.ServeMux for dynamic memory allocation

	api.InitApi(serverMux)

	addr := envConfigs.SrvAddr + ":" + envConfigs.SrvPort

	log.Println("Service running in", addr)

	err := http.ListenAndServe(addr, serverMux)
	if err != nil {
		log.Fatal(err)
	}
}
