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
	if err != nil {
		data := Response{nil, "error"}
		encode(w, data)
		return err
	}
	data := Response{values, "success"}
	encode(w, data)
	return nil
}
func byID(model models.Modeler, w http.ResponseWriter, r *http.Request, app *service.AppServer) error {
	val := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(val)
	err := model.ByID(app.DB, uint(id))
	if err != nil {
		data := Response{nil, err.Error()}
		encode(w, data)
		return err
	}

	data := Response{
		Data:    model,
		Message: "success",
	}
	encode(w, data)
	return nil
}

func encode(w http.ResponseWriter, res Response) {
	if er := json.NewEncoder(w).Encode(res); er != nil {
		log.Println(er)
	}
}
