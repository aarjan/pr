package controller

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/ccdb-api/app/models"
	"github.com/ccdb-api/app/service"
	"github.com/gorilla/mux"
)

func GetClients(w http.ResponseWriter, r *http.Request, app *service.AppServer) error {
	cl := &models.Client{}
	return allData(cl, w, app)
}

func GetSubscriptions(w http.ResponseWriter, r *http.Request, app *service.AppServer) error {
	s := &models.Subscription{}
	return allData(s, w, app)
}

func GetPayments(w http.ResponseWriter, r *http.Request, app *service.AppServer) error {
	s := &models.Payment{}
	return allData(s, w, app)
}

func GetPartners(w http.ResponseWriter, r *http.Request, app *service.AppServer) error {
	s := &models.Partner{}
	return allData(s, w, app)
}

func GetPackages(w http.ResponseWriter, r *http.Request, app *service.AppServer) error {
	s := &models.Package{}
	return allData(s, w, app)
}

func GetSalesPersons(w http.ResponseWriter, r *http.Request, app *service.AppServer) error {
	s := &models.SalesPerson{}
	return allData(s, w, app)
}

func GetSalesPerson(w http.ResponseWriter, r *http.Request, app *service.AppServer) error {
	s := &models.SalesPerson{}
	return byID(s, w, r, app)
}

func GetPackage(w http.ResponseWriter, r *http.Request, app *service.AppServer) error {
	s := &models.Package{}
	return byID(s, w, r, app)
}

func GetPartner(w http.ResponseWriter, r *http.Request, app *service.AppServer) error {
	s := &models.Partner{}
	return byID(s, w, r, app)
}

func GetClient(w http.ResponseWriter, r *http.Request, app *service.AppServer) error {
	val := mux.Vars(r)["id"]
	id, err := strconv.Atoi(val)
	if err != nil {
		encodeErr(w, Response{}, err)
		return err
	}
	cl := &models.Client{}
	err = cl.ByID(app.DB, uint(id))
	if err != nil {
		encodeErr(w, Response{}, err)
		return err
	}
	sp, err := cl.SalesPerson(app.DB)
	if err != sql.ErrNoRows && err != nil {
		encodeErr(w, Response{}, err)
		return err
	}

	data := Response{
		Data: struct {
			Client      *models.Client      `json:"client"`
			SalesPerson *models.SalesPerson `json:"sales_person"`
		}{cl, sp},
		Message: "success",
	}
	encodeErr(w, data, err)
	return nil
}

func GetSubscription(w http.ResponseWriter, r *http.Request, app *service.AppServer) error {
	val := mux.Vars(r)["id"]
	id, err := strconv.Atoi(val)
	if err != nil {
		encodeErr(w, Response{}, err)
		return err
	}
	s := &models.Subscription{}
	err = s.ByID(app.DB, uint(id))
	if err != nil {
		encodeErr(w, Response{}, err)
		return err
	}
	cl, err := s.Client(app.DB)
	if err != nil {
		encodeErr(w, Response{}, err)
		return err
	}
	pack, err := s.Package(app.DB)
	if err != nil {
		encodeErr(w, Response{}, err)
		return err
	}
	part, err := s.Partner(app.DB)
	if err != sql.ErrNoRows && err != nil {
		encodeErr(w, Response{}, err)
		return err
	}
	data := Response{
		Data: struct {
			Subscription *models.Subscription `json:"subscription"`
			Client       *models.Client       `json:"client"`
			Package      *models.Package      `json:"package"`
			Partner      *models.Partner      `json:"partner"`
		}{s, cl, pack, part},
		Message: "success",
	}
	encodeErr(w, data, err)
	return nil
}

func GetPayment(w http.ResponseWriter, r *http.Request, app *service.AppServer) error {
	val := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(val)
	p := &models.Payment{}
	err := p.ByID(app.DB, uint(id))
	if err != nil {
		encodeErr(w, Response{}, err)
		return err
	}
	s, err := p.Subscription(app.DB)
	if err != nil {
		encodeErr(w, Response{}, err)
		return err
	}
	data := Response{
		Data: struct {
			Payment      *models.Payment      `json:"payment"`
			Subscription *models.Subscription `json:"subscription"`
		}{p, s},
		Message: "success",
	}
	encodeErr(w, data, err)
	return nil
}
