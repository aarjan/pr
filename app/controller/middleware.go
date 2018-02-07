package controller

import (
	"database/sql"
	"net/http"

	"github.com/ccdb-api/app/models"
	"github.com/ccdb-api/app/service"
	log "github.com/sirupsen/logrus"
)

type AppHandler struct {
	*service.AppServer
	HandlerFunc func(w http.ResponseWriter, r *http.Request, app *service.AppServer) error
}

func (app AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers",
		"Accept, Accept-Version, Api-Version, Content-Length, Content-MD5, Content-Type, Content-Disposition, Date, Origin, X-Dragon-Law-Dragonball, X-Dragon-Law-Username, X-Dragon-Law-Signer, X-Dragon-Law-Token, X-Dragon-Law-API-Version, X-Requested-With")
	w.Header().Set("Access-Control-Expose-Headers",
		"Api-Version, Cache-Control, Content-Length, Content-MD5, Content-Type, Content-Disposition, Connection-Status, Date, Expires, Origin, Pragma, Request-ID, Response-Time, X-Chopped, X-Dragon-Law-Dragonball, X-Dragon-Law-API-Version, X-Dragon-Law-Username, X-Dragon-Law-Signer, X-Dragon-Law-Token",
	)
	w.Header().Set("Accept-Encoding", "gzip,deflate,br")

	// Run the respective handler
	err := app.HandlerFunc(w, r, app.AppServer)
	if err != nil {
		log.WithFields(log.Fields{
			"Error":      err.Error(),
			"Method":     r.Method,
			"CurrentURL": r.URL.Path,
		}).Warn()
	}
}

func AuthMiddleware(h http.Handler) http.Handler {
	appHandler := h.(AppHandler)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		log.WithFields(log.Fields{
			"Referrer":   r.Referer(),
			"CurrentURL": r.URL.Path,
			"Method":     r.Method,
		}).Info("Auth Handler Logging!!")

		// Auth check
		accessToken, ok := r.Header["X-Dragon-Law-Dragonball"]
		if !ok {
			log.Info("Auth Handler redirecting", "no accesstoken")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		username, ok := r.Header["X-Dragon-Law-Username"]
		if !ok {
			log.Info("Auth Handler redirecting", "no username")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		u, err := models.UserByUsername(appHandler.DB, username[0])
		if err != nil {
			if err == sql.ErrNoRows {
				log.Info("Auth Handler redirecting", "no username in db")
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			http.Error(w, "Server Error", 505)
			log.WithFields(log.Fields{
				"Error": err.Error(),
			}).Warn()
			return
		}
		if accessToken[0] != u.AccessToken.String {
			log.Info(accessToken[0], u.AccessToken)
			log.Info("Auth Handler redirecting", "accesstoken not matched")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		appHandler.ServeHTTP(w, r)
	})
}
