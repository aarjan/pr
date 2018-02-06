package controller

import (
	"encoding/json"
	"net/http"

	"github.com/ccdb-api/app/models"
	"github.com/ccdb-api/app/service"
)

func CreateClient(w http.ResponseWriter, r *http.Request, app *service.AppServer) error {
	m := &models.Client{}
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		encodeErr(w, Response{}, err)
		return err
	}
	err = m.Insert(app.DB)
	res := Response{m, "create success"}
	encodeErr(w, res, err)
	return err
}

func CreateSalesPerson(w http.ResponseWriter, r *http.Request, app *service.AppServer) error {
	m := &models.SalesPerson{}
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		encodeErr(w, Response{}, err)
		return err
	}
	err = m.Insert(app.DB)
	res := Response{m, "create success"}
	encodeErr(w, res, err)
	return err
}

func CreatePackage(w http.ResponseWriter, r *http.Request, app *service.AppServer) error {
	m := &models.Package{}
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		encodeErr(w, Response{}, err)
		return err
	}
	err = m.Insert(app.DB)
	res := Response{m, "create success"}
	encodeErr(w, res, err)
	return err
}

func CreatePartner(w http.ResponseWriter, r *http.Request, app *service.AppServer) error {
	m := &models.Partner{}
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		encodeErr(w, Response{}, err)
		return err
	}
	err = m.Insert(app.DB)
	res := Response{m, "create success"}
	encodeErr(w, res, err)
	return err
}

func CreatePayment(w http.ResponseWriter, r *http.Request, app *service.AppServer) error {
	m := &models.Payment{}
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		encodeErr(w, Response{}, err)
		return err
	}
	err = m.Insert(app.DB)
	res := Response{m, "create success"}
	encodeErr(w, res, err)
	return err
}

func CreateSubscription(w http.ResponseWriter, r *http.Request, app *service.AppServer) error {
	m := &models.Subscription{}
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		encodeErr(w, Response{}, err)
		return err
	}
	err = m.Insert(app.DB)
	res := Response{m, "create success"}
	encodeErr(w, res, err)
	return err
}
