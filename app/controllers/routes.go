package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (server *Server) initializeroutes() {
	server.Router = mux.NewRouter()
	server.Router.HandleFunc("/", server.Home).Methods("GET")
	server.Router.HandleFunc("/product", server.Product).Methods("GET")

	staticFileDirectory := http.Dir("./assets/")
	staticFileHandler := http.StripPrefix("/public/", http.FileServer(staticFileDirectory))
	server.Router.PathPrefix("/public/").Handler(staticFileHandler).Methods("GET")
}
