package controller

import (
	"log"
	"net/http"

	"github.com/ccdb-api/app/service"
)

type AppHandler struct {
	*service.AppServer
	HandlerFunc func(w http.ResponseWriter, r *http.Request, app *service.AppServer) error
}

func (app AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("Method", r.Method)
	log.Println("Route", r.URL.Path)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	err := app.HandlerFunc(w, r, app.AppServer)
	if err != nil {
		log.Println("server error", err)
	}
}
