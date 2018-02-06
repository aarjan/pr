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

	m.Handle("/login", AppHandler{app, Login}).Methods("POST")

	// GET
	m.Handle("/clients", AuthMiddleware(AppHandler{app, GetClients})).Methods("GET")
	m.Handle("/subscriptions", AppHandler{app, GetSubscriptions}).Methods("GET")
	m.Handle("/payments", AppHandler{app, GetPayments}).Methods("GET")
	m.Handle("/salespersons", AppHandler{app, GetSalesPersons}).Methods("GET")
	m.Handle("/packages", AppHandler{app, GetPackages}).Methods("GET")
	m.Handle("/partners", AppHandler{app, GetPartners}).Methods("GET")

	m.Handle("/client/{id}", AppHandler{app, GetClient}).Methods("GET")
	m.Handle("/subscription/{id}", AppHandler{app, GetSubscription}).Methods("GET")
	m.Handle("/payment/{id}", AppHandler{app, GetPayment}).Methods("GET")
	m.Handle("/partner/{id}", AppHandler{app, GetPartner}).Methods("GET")
	m.Handle("/package/{id}", AppHandler{app, GetPackage}).Methods("GET")
	m.Handle("/salesperson/{id}", AppHandler{app, GetSalesPerson}).Methods("GET")

	// Delete
	m.Handle("/client/{id}", AppHandler{app, DeleteClient}).Methods("DELETE")
	m.Handle("/partner/{id}", AppHandler{app, DeletePartner}).Methods("DELETE")
	m.Handle("/package/{id}", AppHandler{app, DeletePackage}).Methods("DELETE")
	m.Handle("/payment/{id}", AppHandler{app, DeletePayment}).Methods("DELETE")
	m.Handle("/salesperson/{id}", AppHandler{app, DeleteSalesPerson}).Methods("DELETE")
	m.Handle("/subscription/{id}", AppHandler{app, DeleteSubscription}).Methods("DELETE")

	// Create
	m.Handle("/client", AppHandler{app, CreateClient}).Methods("POST")
	m.Handle("/salesperson", AppHandler{app, CreateSalesPerson}).Methods("POST")
	m.Handle("/partner", AppHandler{app, CreatePartner}).Methods("POST")
	m.Handle("/package", AppHandler{app, CreatePackage}).Methods("POST")
	m.Handle("/subscription", AppHandler{app, CreateSubscription}).Methods("POST")
	m.Handle("/payment", AppHandler{app, CreatePayment}).Methods("POST")

	// Update
	m.Handle("/client/{id}", AppHandler{app, UpdateClient}).Methods("PUT")
	m.Handle("/salesperson/{id}", AppHandler{app, UpdateSalesPerson}).Methods("PUT")
	m.Handle("/partner/{id}", AppHandler{app, UpdatePartner}).Methods("PUT")
	m.Handle("/package/{id}", AppHandler{app, UpdatePackage}).Methods("PUT")
	m.Handle("/subscription/{id}", AppHandler{app, UpdateSubscription}).Methods("PUT")
	m.Handle("/payment/{id}", AppHandler{app, UpdatePayment}).Methods("PUT")

	return m
}
