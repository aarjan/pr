package controller

import (
	"net/http"

	"github.com/ccdb-api/app/models"
	"github.com/ccdb-api/app/service"
)

type Response struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func Clients(w http.ResponseWriter, r *http.Request, app *service.AppServer) error {
	cl := models.Client{}
	return allData(cl, w, r, app)
}

func Subscriptions(w http.ResponseWriter, r *http.Request, app *service.AppServer) error {
	s := models.Subscription{}
	return allData(s, w, r, app)
}

func Payments(w http.ResponseWriter, r *http.Request, app *service.AppServer) error {
	s := models.Payment{}
	return allData(s, w, r, app)
}

func Partners(w http.ResponseWriter, r *http.Request, app *service.AppServer) error {
	s := models.Partner{}
	return allData(s, w, r, app)
}

func Packages(w http.ResponseWriter, r *http.Request, app *service.AppServer) error {
	s := models.Package{}
	return allData(s, w, r, app)
}

func SalesPersons(w http.ResponseWriter, r *http.Request, app *service.AppServer) error {
	s := models.SalesPerson{}
	return allData(s, w, r, app)
}

// func Client(w http.ResponseWriter, r *http.Request, app *service.AppServer) error {
// 	val := mux.Vars(r)["id"]
// 	id, _ := strconv.Atoi(val)
// 	cl := models.Client{}
// 	client, err := cl.ByID(app.DB, uint(id))
// 	if err != nil {
// 		data := Response{nil, err.Error()}
// 		encode(w, data)
// 		return err
// 	}
// 	sp, err := client.(models.Client).SalesPerson(app.DB)
// 	if err != nil {
// 		data := Response{nil, err.Error()}
// 		encode(w, data)
// 		return err
// 	}

// 	data := Response{
// 		Data: struct {
// 			Client      *models.Client      `json:"client"`
// 			SalesPerson *models.SalesPerson `json:"sales_person"`
// 		}{client, sp},
// 		Message: "success",
// 	}
// 	encode(w, data)
// 	return nil
// }

func SalesPerson(w http.ResponseWriter, r *http.Request, app *service.AppServer) error {
	s := models.SalesPerson{}
	return byID(s, w, r, app)
}

func Package(w http.ResponseWriter, r *http.Request, app *service.AppServer) error {
	s := models.Package{}
	return byID(s, w, r, app)
}

func Partner(w http.ResponseWriter, r *http.Request, app *service.AppServer) error {
	s := models.Partner{}
	return byID(s, w, r, app)
}
