package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/ccdb-api/app/models"
	"github.com/ccdb-api/app/service"
)

type Response struct {
	Data  interface{} `json:"data"`
	Error string      `json:"error"`
}

func Clients(w http.ResponseWriter, r *http.Request, app *service.AppServer) error {
	clients, err := models.Clients(app.DB)
	if err != nil {
		data := Response{
			Data:  nil,
			Error: err.Error(),
		}
		encode(w, data)
		return err
	}
	data := Response{
		Data:  clients,
		Error: "",
	}
	encode(w, data)
	return nil
}

func Subscriptions(w http.ResponseWriter, r *http.Request, app *service.AppServer) error {
	s := models.Subscription{}
	return handling(s, w, r, app)
}

func handling(model models.Modeler, w http.ResponseWriter, r *http.Request, app *service.AppServer) error {
	subscriptions, err := model.All(app.DB)
	if err != nil {
		data := Response{
			Data:  nil,
			Error: err.Error(),
		}
		encode(w, data)
		return err
	}
	data := Response{
		Data:  subscriptions,
		Error: "",
	}
	encode(w, data)
	return nil
}

func encode(w http.ResponseWriter, res Response) {
	if er := json.NewEncoder(w).Encode(res); er != nil {
		log.Println(er)
	}
}

func Client(w http.ResponseWriter, r *http.Request, app *service.AppServer) error {
	val := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(val)
	client, err := models.ClientByID(app.DB, uint(id))
	if err != nil {
		data := Response{
			Data:  nil,
			Error: err.Error(),
		}
		encode(w, data)
		return err
	}
	sp, err := client.SalesPerson(app.DB)
	if err != nil {
		data := Response{
			Data:  nil,
			Error: err.Error(),
		}
		encode(w, data)
		return err
	}

	data := Response{
		Data: struct {
			*models.Client
			*models.SalesPerson
		}{client, sp},
		Error: "",
	}
	encode(w, data)
	return nil
}
