package controller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/aarjan/blog/app/shared"
	"github.com/ccdb-api/app/models"
	"github.com/ccdb-api/app/service"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

func Login(w http.ResponseWriter, r *http.Request, app *service.AppServer) error {
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusNoContent)
		return nil
	}
	req := struct {
		Dash     bool   `json:"dash"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return err
	}
	name := req.Email
	pass := req.Password

	// user verification
	user, err := models.VerifyUser(name, pass, app.DB)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return err
	}
	user.LastSeen = time.Now()
	user.Update(app.DB)
	w.Header().Set("X-Dragon-Law-Dragonball", user.AccessToken.String)
	w.Header().Set("X-Dragon-Law-Username", user.Username)
	res := Response{nil, "success"}
	encodeErr(w, res, nil)
	return err
}

func Register(w http.ResponseWriter, r *http.Request, app *service.AppServer) error {
	m := make(map[string]string)
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		encodeErr(w, Response{}, err)
		return errors.Wrap(err, "encoding error")
	}
	name := m["email"]
	pass := m["password"]

	// verify username
	user, err := models.UserByUsername(app.DB, name)
	if user != nil {
		err := errors.New("Username already exists")
		encodeErr(w, Response{}, err)
		return err
	}

	accessToken := shared.GenerateAccessToken(name)
	password, _ := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)

	user = &models.User{
		Username:    name,
		Password:    string(password),
		AccessToken: models.NullString{string(accessToken), true},
		LoginTime:   time.Now(),
		LastSeen:    time.Now(),
	}

	// save record to database
	err = user.Insert(app.DB)
	w.Header().Set("X-Dragon-Law-Dragonball", user.AccessToken.String)
	w.Header().Set("X-Dragon-Law-Username", user.Username)
	res := Response{nil, "success"}
	encodeErr(w, res, err)
	return err
}
