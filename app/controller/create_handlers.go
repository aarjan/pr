package controller

import (
	"net/http"

	"github.com/ccdb-api/app/models"
	"github.com/ccdb-api/app/service"
)

func CreateClient(w http.ResponseWriter, r *http.Request, app *service.AppServer) error {
	m := &models.Client{}
	return createModel(m, w, r, app)
}

func CreateSalesPerson(w http.ResponseWriter, r *http.Request, app *service.AppServer) error {
	m := &models.SalesPerson{}
	return createModel(m, w, r, app)
}

func CreatePackage(w http.ResponseWriter, r *http.Request, app *service.AppServer) error {
	m := &models.Package{}
	return createModel(m, w, r, app)
}

func CreatePartner(w http.ResponseWriter, r *http.Request, app *service.AppServer) error {
	m := &models.Partner{}
	return createModel(m, w, r, app)
}

func CreatePayment(w http.ResponseWriter, r *http.Request, app *service.AppServer) error {
	m := &models.Payment{}
	return createModel(m, w, r, app)
}

func CreateSubscription(w http.ResponseWriter, r *http.Request, app *service.AppServer) error {
	m := &models.Subscription{}
	return createModel(m, w, r, app)
}
