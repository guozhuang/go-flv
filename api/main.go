package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

	router.GET("/user", CreateUser)

	router.GET("/user/:user_name", Login)

	return router
}

func main() {
	r := RegisterHandlers()

	http.ListenAndServe(":80", r)
}
