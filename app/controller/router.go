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
	m.Handle("/subscriptions", AppHandler{app, Subscriptions})
	m.Handle("/payments", AppHandler{app, Payments})
	m.Handle("/salespersons", AppHandler{app, SalesPersons})
	m.Handle("/packages", AppHandler{app, Packages})
	m.Handle("/partners", AppHandler{app, Partners})

	m.Handle("/partner/{id}", AppHandler{app, Partner})
	m.Handle("/package/{id}", AppHandler{app, Package})
	m.Handle("/salesperson/{id}", AppHandler{app, SalesPerson})
	return m
}
