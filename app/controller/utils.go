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

func allData(model models.Modeler, w http.ResponseWriter, r *http.Request, app *service.AppServer) error {
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
