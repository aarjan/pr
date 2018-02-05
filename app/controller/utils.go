package controller

import (
	"encoding/json"
	"log"
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

func createModel(model models.Modeler, w http.ResponseWriter, r *http.Request, app *service.AppServer) error {
	var err error
	err = json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		res := Response{nil, err.Error()}
		encode(w, res)
		return err
	}
	switch m := model.(type) {
	case *models.Client:
		err = m.Insert(app.DB)
		model = m

	case *models.SalesPerson:
		err = m.Insert(app.DB)
		model = m

	case *models.Partner:
		err = m.Insert(app.DB)
		model = m

	case *models.Package:
		err = m.Insert(app.DB)
		model = m

	case *models.Subscription:
		err = m.Insert(app.DB)
		model = m

	case *models.Payment:
		err = m.Insert(app.DB)
		model = m
	}
	res := Response{model, "success"}
	encodeErr(w, res, err)
	return err
}

func deleteModel(model models.Modeler, w http.ResponseWriter, r *http.Request, app *service.AppServer) error {
	var err error
	val := mux.Vars(r)["id"]
	id, err := strconv.Atoi(val)
	if err != nil {
		data := Response{nil, err.Error()}
		encode(w, data)
		return err
	}
	switch m := model.(type) {
	case *models.Client:
		m.ID = uint(id)
		err = m.Delete(app.DB)
		model = m
	case *models.SalesPerson:
		m.ID = uint(id)
		err = m.Delete(app.DB)
		model = m
	case *models.Partner:
		m.ID = uint(id)
		err = m.Delete(app.DB)
		model = m
	case *models.Package:
		m.ID = uint(id)
		err = m.Delete(app.DB)
		model = m
	case *models.Subscription:
		m.ID = uint(id)
		err = m.Delete(app.DB)
		model = m
	case *models.Payment:
		m.ID = uint(id)
		err = m.Delete(app.DB)
		model = m
	}
	data := Response{model, "success"}
	encodeErr(w, data, err)
	return nil
}

func allData(model models.Modeler, w http.ResponseWriter, app *service.AppServer) error {
	values, err := model.All(app.DB)
	data := Response{values, "success"}
	encodeErr(w, data, err)
	return err
}

func byID(model models.Modeler, w http.ResponseWriter, r *http.Request, app *service.AppServer) error {
	val := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(val)
	err := model.ByID(app.DB, uint(id))

	var d interface{}
	switch model.(type) {
	case *models.SalesPerson:
		d = struct {
			SalesPerson *models.SalesPerson `json:"salesperson"`
		}{model.(*models.SalesPerson)}
	case *models.Package:
		d = struct {
			Package *models.Package `json:"package"`
		}{model.(*models.Package)}
	case *models.Partner:
		d = struct {
			Partner *models.Partner `json:"partner"`
		}{model.(*models.Partner)}
	}
	data := Response{
		Data:    d,
		Message: "success",
	}
	encodeErr(w, data, err)
	return err
}

func encodeErr(w http.ResponseWriter, res Response, err error) {
	if err != nil {
		res = Response{nil, err.Error()}
	}
	if er := json.NewEncoder(w).Encode(res); er != nil {
		log.Println(er)
	}
}

func encode(w http.ResponseWriter, res Response) {
	if er := json.NewEncoder(w).Encode(res); er != nil {
		log.Println(er)
	}
}
