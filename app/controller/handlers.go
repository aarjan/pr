package controller

import (
	"errors"
	"net/http"

	"github.com/aarjan/blog/app/shared"
	"github.com/ccdb-api/app/models"
	"github.com/ccdb-api/app/service"
	"golang.org/x/crypto/bcrypt"
)

func Login(w http.ResponseWriter, r *http.Request, app *service.AppServer) error {
	r.ParseForm()
	name := r.FormValue("username")
	pass := r.FormValue("password")

	// user verification
	user, err := models.VerifyUser(name, pass, app.DB)
	res := Response{user, "success"}
	encodeErr(w, res, err)
	return err
}

func Register(w http.ResponseWriter, r *http.Request, app *service.AppServer) error {

	r.ParseForm()
	name := r.FormValue("username")
	pass := r.FormValue("password")

	// verify username
	user, err := models.UserByUsername(app.DB, name)
	if user.ID != 0 {
		err := errors.New("Username already exists")
		encodeErr(w, Response{}, err)
		return err
	}

	accessToken := shared.GenerateAccessToken(name)
	password, _ := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)

	user.AccessToken = models.NullString{string(accessToken), true}
	user.Password = string(password)

	// save record to database
	err = user.Insert(app.DB)
	res := Response{user, "success"}
	encodeErr(w, res, err)
	return err
}
