package api

import "github.com/gorilla/mux"

func InitRoute() (r *mux.Router) {
	api := NewHandler()
	r = mux.NewRouter()

	r.HandleFunc("/menus", api.GetMenu).Methods("GET")
	r.HandleFunc("/chef", api.AddChef).Methods("POST")
	r.HandleFunc("/orders", api.AddOrder).Methods("POST")

	return r
}
