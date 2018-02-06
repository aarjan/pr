package controller

import (
	"database/sql"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/ccdb-api/app/models"
	"github.com/ccdb-api/app/service"
)

type AppHandler struct {
	*service.AppServer
	HandlerFunc func(w http.ResponseWriter, r *http.Request, app *service.AppServer) error
}

func (app AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("Method", r.Method)
	log.Println("Route", r.URL.Path)

	w.Header().Set("Access-Control-Allow-Headers",
		"Accept,Accept-Version, Api-Version,"+
			"Content-Length, Content-MD5,Content-Type,"+
			"Content-Disposition, Date, Origin, X-Dragon-Law-Dragonball,"+
			"X-Dragon-Law-Username, X-Dragon-Law-Signer, X-Dragon-Law-Token,"+
			"X-Dragon-Law-API-Version, X-Requested-With")

	w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Content-Type", "application/json")

	// Run the respective handler
	err := app.HandlerFunc(w, r, app.AppServer)
	if err != nil {
		log.Println("server error", err)
	}
}

func AuthMiddleware(h http.Handler) http.Handler {
	appHandler := h.(AppHandler)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Auth check
		accessToken, ok := r.Header["X-Dragon-Law-Dragonball"]
		if !ok {
			log.Println("Auth Handler redirecting")
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		username, ok := r.Header["X-Dragon-Law-Username"]
		if !ok {
			log.Println("Auth Handler redirecting")
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		u, err := models.UserByUsername(appHandler.DB, username[0])
		if err != nil {
			if err == sql.ErrNoRows {
				log.Println("Auth Handler redirecting")
				http.Redirect(w, r, "/", http.StatusFound)
				return
			}
			http.Error(w, "", 505)
			log.Println(err)
			return
		}
		err = bcrypt.CompareHashAndPassword([]byte(accessToken[0]), []byte(u.AccessToken.String))
		if err != nil {
			log.Println("Auth Handler redirecting")
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}

		appHandler.ServeHTTP(w, r)
	})
}
