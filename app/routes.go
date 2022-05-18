package app

import (
	"github.com/Jehanv60/gotoko/app/controllers"
	"github.com/gorilla/mux"
)

func (server *Server) initializeroutes() {
	server.Router = mux.NewRouter()
	server.Router.HandleFunc("/", controllers.Home).Methods("GET")
}
