package controller

import (
	"encoding/json"
	"net/http"

	"github.com/ccdb-api/app/models"
	"github.com/ccdb-api/app/service"
)

func UpdateClient(w http.ResponseWriter, r *http.Request, app *service.AppServer) (err error) {
	c := &models.Client{}
	err = json.NewDecoder(r.Body).Decode(c)
	if err != nil {
		encodeErr(w, Response{}, err)
		return err
	}
	err = c.Update(app.DB)
	res := Response{c, "success"}
	encodeErr(w, res, err)
	return err
}

func UpdateSalesPerson(w http.ResponseWriter, r *http.Request, app *service.AppServer) (err error) {
	c := &models.SalesPerson{}
	err = json.NewDecoder(r.Body).Decode(c)
	if err != nil {
		encodeErr(w, Response{}, err)
		return err
	}
	err = c.Update(app.DB)
	res := Response{c, "success"}
	encodeErr(w, res, err)
	return err
}
func UpdatePackage(w http.ResponseWriter, r *http.Request, app *service.AppServer) (err error) {
	c := &models.Package{}
	err = json.NewDecoder(r.Body).Decode(c)
	if err != nil {
		encodeErr(w, Response{}, err)
		return err
	}
	err = c.Update(app.DB)
	res := Response{c, "success"}
	encodeErr(w, res, err)
	return err
}
func UpdatePartner(w http.ResponseWriter, r *http.Request, app *service.AppServer) (err error) {
	c := &models.Partner{}
	err = json.NewDecoder(r.Body).Decode(c)
	if err != nil {
		encodeErr(w, Response{}, err)
		return err
	}
	err = c.Update(app.DB)
	res := Response{c, "success"}
	encodeErr(w, res, err)
	return err
}
func UpdateSubscription(w http.ResponseWriter, r *http.Request, app *service.AppServer) (err error) {
	c := &models.Subscription{}
	err = json.NewDecoder(r.Body).Decode(c)
	if err != nil {
		encodeErr(w, Response{}, err)
		return err
	}
	err = c.Update(app.DB)
	res := Response{c, "success"}
	encodeErr(w, res, err)
	return err
}
func UpdatePayment(w http.ResponseWriter, r *http.Request, app *service.AppServer) (err error) {
	c := &models.Payment{}
	err = json.NewDecoder(r.Body).Decode(c)
	if err != nil {
		encodeErr(w, Response{}, err)
		return err
	}
	err = c.Update(app.DB)
	res := Response{c, "success"}
	encodeErr(w, res, err)
	return err
}
