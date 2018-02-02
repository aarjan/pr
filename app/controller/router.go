package controller

import (
	"fmt"
	"net/http"

	"github.com/ccdb-api/app/service"
	"github.com/gorilla/mux"
)

func Handler(app *service.AppServer) *mux.Router {
	m := mux.NewRouter()
	m.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Welcome to home directory")
	})

	m.Handle("/clients", AppHandler{app, Clients})
	return m
}
