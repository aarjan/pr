package controller

import (
	"net/http"
	"strconv"

	"github.com/ccdb-api/app/models"
	"github.com/ccdb-api/app/service"
	"github.com/gorilla/mux"
)

type Response struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func Clients(w http.ResponseWriter, r *http.Request, app *service.AppServer) error {
	cl := &models.Client{}
	return allData(cl, w, r, app)
}

func Subscriptions(w http.ResponseWriter, r *http.Request, app *service.AppServer) error {
	s := &models.Subscription{}
	return allData(s, w, r, app)
}

func Payments(w http.ResponseWriter, r *http.Request, app *service.AppServer) error {
	s := &models.Payment{}
	return allData(s, w, r, app)
}

func Partners(w http.ResponseWriter, r *http.Request, app *service.AppServer) error {
	s := &models.Partner{}
	return allData(s, w, r, app)
}

func Packages(w http.ResponseWriter, r *http.Request, app *service.AppServer) error {
	s := &models.Package{}
	return allData(s, w, r, app)
}

func SalesPersons(w http.ResponseWriter, r *http.Request, app *service.AppServer) error {
	s := &models.SalesPerson{}
	return allData(s, w, r, app)
}

func SalesPerson(w http.ResponseWriter, r *http.Request, app *service.AppServer) error {
	s := &models.SalesPerson{}
	return byID(s, w, r, app)
}

func Package(w http.ResponseWriter, r *http.Request, app *service.AppServer) error {
	s := &models.Package{}
	return byID(s, w, r, app)
}

func Partner(w http.ResponseWriter, r *http.Request, app *service.AppServer) error {
	s := &models.Partner{}
	return byID(s, w, r, app)
}

func Client(w http.ResponseWriter, r *http.Request, app *service.AppServer) error {
	val := mux.Vars(r)["id"]
	id, err := strconv.Atoi(val)
	if err != nil {
		data := Response{nil, err.Error()}
		encode(w, data)
		return err
	}
	cl := &models.Client{}
	err = cl.ByID(app.DB, uint(id))
	if err != nil {
		data := Response{nil, err.Error()}
		encode(w, data)
		return err
	}
	sp, err := cl.SalesPerson(app.DB)
	if err != nil {
		data := Response{nil, err.Error()}
		encode(w, data)
		return err
	}

	data := Response{
		Data: struct {
			Client      *models.Client      `json:"client"`
			SalesPerson *models.SalesPerson `json:"sales_person"`
		}{cl, sp},
		Message: "success",
	}
	encode(w, data)
	return nil
}

func Subscription(w http.ResponseWriter, r *http.Request, app *service.AppServer) error {
	val := mux.Vars(r)["id"]
	id, err := strconv.Atoi(val)
	if err != nil {
		data := Response{nil, err.Error()}
		encode(w, data)
		return err
	}
	s := &models.Subscription{}
	err = s.ByID(app.DB, uint(id))
	if err != nil {
		data := Response{nil, err.Error()}
		encode(w, data)
		return err
	}
	cl, err := s.Client(app.DB)
	if err != nil {
		data := Response{nil, err.Error()}
		encode(w, data)
		return err
	}
	pack, err := s.Package(app.DB)
	if err != nil {
		data := Response{nil, err.Error()}
		encode(w, data)
		return err
	}
	part, err := s.Partner(app.DB)
	if err != nil {
		data := Response{nil, err.Error()}
		encode(w, data)
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
	encode(w, data)
	return nil
}

func Payment(w http.ResponseWriter, r *http.Request, app *service.AppServer) error {
	val := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(val)
	p := &models.Payment{}
	err := p.ByID(app.DB, uint(id))
	if err != nil {
		data := Response{nil, err.Error()}
		encode(w, data)
		return err
	}
	s, err := p.Subscription(app.DB)
	if err != nil {
		data := Response{nil, err.Error()}
		encode(w, data)
		return err
	}
	data := Response{
		Data: struct {
			Payment      *models.Payment      `json:"payment"`
			Subscription *models.Subscription `json:"subscription"`
		}{p, s},
		Message: "success",
	}
	encode(w, data)
	return nil
}
