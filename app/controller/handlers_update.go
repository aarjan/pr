package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ccdb-api/app/models"
	"github.com/ccdb-api/app/service"
	"github.com/gorilla/mux"
)

func UpdateClient(w http.ResponseWriter, r *http.Request, app *service.AppServer) error {
	val := mux.Vars(r)["id"]
	id, err := strconv.Atoi(val)
	if err != nil {
		encodeErr(w, Response{}, err)
		return err
	}
	c := &models.Client{ID: uint(id)}
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

func UpdateSalesPerson(w http.ResponseWriter, r *http.Request, app *service.AppServer) error {
	val := mux.Vars(r)["id"]
	id, err := strconv.Atoi(val)
	if err != nil {
		encodeErr(w, Response{}, err)
		return err
	}
	c := &models.SalesPerson{ID: uint(id)}
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
func UpdatePackage(w http.ResponseWriter, r *http.Request, app *service.AppServer) error {
	val := mux.Vars(r)["id"]
	id, err := strconv.Atoi(val)
	if err != nil {
		encodeErr(w, Response{}, err)
		return err
	}
	c := &models.Package{ID: uint(id)}
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
func UpdatePartner(w http.ResponseWriter, r *http.Request, app *service.AppServer) error {
	val := mux.Vars(r)["id"]
	id, err := strconv.Atoi(val)
	if err != nil {
		encodeErr(w, Response{}, err)
		return err
	}
	c := &models.Partner{ID: uint(id)}
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
func UpdateSubscription(w http.ResponseWriter, r *http.Request, app *service.AppServer) error {
	val := mux.Vars(r)["id"]
	id, err := strconv.Atoi(val)
	if err != nil {
		encodeErr(w, Response{}, err)
		return err
	}
	c := &models.Subscription{ID: uint(id)}
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
func UpdatePayment(w http.ResponseWriter, r *http.Request, app *service.AppServer) error {
	val := mux.Vars(r)["id"]
	id, err := strconv.Atoi(val)
	if err != nil {
		encodeErr(w, Response{}, err)
		return err
	}
	c := &models.Payment{ID: uint(id)}
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
