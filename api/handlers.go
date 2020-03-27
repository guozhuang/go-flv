package main

import (
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
	"dbops"
)

func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	dbops.AddUserCredential()
}

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	uname := p.ByName("user_name")

	io.WriteString(w, uname)
}
