package handlers

import (
	"fmt"
	"net/http"
)

func Login(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Login")
}

func Regist(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Regist")
}
