package api

import (
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func InitRoute() (r *mux.Router) {
	rLog := log.WithField("api", "router")
	api := NewHandler()
	r = mux.NewRouter()

	r.HandleFunc("/{url}", api.RedirectURL).Methods("GET")

	r.HandleFunc("/api/shorten", api.ShortURL).Methods("POST")
	r.HandleFunc("/api/list", api.GetListURL).Methods("GET")

	rLog.Info("route initialize")
	return r
}
