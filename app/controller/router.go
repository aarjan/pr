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

	m.Handle("/auth/login", AppHandler{app, Login}).Methods("POST", "OPTIONS")
	m.Handle("/register", AppHandler{app, Register}).Methods("POST")

	// GET
	m.Handle("/clients", AuthMiddleware(AppHandler{app, GetClients})).Methods("GET")
	m.Handle("/subscriptions", AuthMiddleware(AppHandler{app, GetSubscriptions})).Methods("GET")
	m.Handle("/payments", AuthMiddleware(AppHandler{app, GetPayments})).Methods("GET")
	m.Handle("/salespersons", AuthMiddleware(AppHandler{app, GetSalesPersons})).Methods("GET")
	m.Handle("/packages", AuthMiddleware(AppHandler{app, GetPackages})).Methods("GET")
	m.Handle("/partners", AuthMiddleware(AppHandler{app, GetPartners})).Methods("GET")

	m.Handle("/client/{id}", AuthMiddleware(AppHandler{app, GetClient})).Methods("GET")
	m.Handle("/subscription/{id}", AuthMiddleware(AppHandler{app, GetSubscription})).Methods("GET")
	m.Handle("/payment/{id}", AuthMiddleware(AppHandler{app, GetPayment})).Methods("GET")
	m.Handle("/partner/{id}", AuthMiddleware(AppHandler{app, GetPartner})).Methods("GET")
	m.Handle("/package/{id}", AuthMiddleware(AppHandler{app, GetPackage})).Methods("GET")
	m.Handle("/salesperson/{id}", AuthMiddleware(AppHandler{app, GetSalesPerson})).Methods("GET")

	// Delete
	m.Handle("/client/{id}", AuthMiddleware(AppHandler{app, DeleteClient})).Methods("DELETE")
	m.Handle("/partner/{id}", AuthMiddleware(AppHandler{app, DeletePartner})).Methods("DELETE")
	m.Handle("/package/{id}", AuthMiddleware(AppHandler{app, DeletePackage})).Methods("DELETE")
	m.Handle("/payment/{id}", AuthMiddleware(AppHandler{app, DeletePayment})).Methods("DELETE")
	m.Handle("/salesperson/{id}", AuthMiddleware(AppHandler{app, DeleteSalesPerson})).Methods("DELETE")
	m.Handle("/subscription/{id}", AuthMiddleware(AppHandler{app, DeleteSubscription})).Methods("DELETE")

	// Create
	m.Handle("/client", AuthMiddleware(AppHandler{app, CreateClient})).Methods("POST")
	m.Handle("/salesperson", AuthMiddleware(AppHandler{app, CreateSalesPerson})).Methods("POST")
	m.Handle("/partner", AuthMiddleware(AppHandler{app, CreatePartner})).Methods("POST")
	m.Handle("/package", AuthMiddleware(AppHandler{app, CreatePackage})).Methods("POST")
	m.Handle("/subscription", AuthMiddleware(AppHandler{app, CreateSubscription})).Methods("POST")
	m.Handle("/payment", AuthMiddleware(AppHandler{app, CreatePayment})).Methods("POST")

	// Update
	m.Handle("/client/{id}", AuthMiddleware(AppHandler{app, UpdateClient})).Methods("PUT")
	m.Handle("/salesperson/{id}", AuthMiddleware(AppHandler{app, UpdateSalesPerson})).Methods("PUT")
	m.Handle("/partner/{id}", AuthMiddleware(AppHandler{app, UpdatePartner})).Methods("PUT")
	m.Handle("/package/{id}", AuthMiddleware(AppHandler{app, UpdatePackage})).Methods("PUT")
	m.Handle("/subscription/{id}", AuthMiddleware(AppHandler{app, UpdateSubscription})).Methods("PUT")
	m.Handle("/payment/{id}", AuthMiddleware(AppHandler{app, UpdatePayment})).Methods("PUT")

	return m
}
