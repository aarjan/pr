package controller

import (
	"net/http"

	"github.com/ccdb-api/app/models"
	"github.com/ccdb-api/app/service"
)

func DeleteClient(w http.ResponseWriter, r *http.Request, app *service.AppServer) error {
	cl := &models.Client{}
	return deleteModel(cl, w, r, app)
}

func DeleteSalesPerson(w http.ResponseWriter, r *http.Request, app *service.AppServer) error {
	cl := &models.SalesPerson{}
	return deleteModel(cl, w, r, app)
}

func DeletePartner(w http.ResponseWriter, r *http.Request, app *service.AppServer) error {
	cl := &models.Partner{}
	return deleteModel(cl, w, r, app)
}

func DeletePackage(w http.ResponseWriter, r *http.Request, app *service.AppServer) error {
	cl := &models.Package{}
	return deleteModel(cl, w, r, app)
}

func DeleteSubscription(w http.ResponseWriter, r *http.Request, app *service.AppServer) error {
	cl := &models.Subscription{}
	return deleteModel(cl, w, r, app)
}

func DeletePayment(w http.ResponseWriter, r *http.Request, app *service.AppServer) error {
	cl := &models.Payment{}
	return deleteModel(cl, w, r, app)
}
