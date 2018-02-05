package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ccdb-api/app/service"
)

func UpdateClient(w http.ResponseWriter, r *http.Request, app *service.AppServer) error {
	var v interface{}
	fmt.Println(json.NewDecoder(r.Body).Decode(&v), v)
	return nil
}
