package app

import "github.com/Jehanv60/gotoko/app/controllers"

func (server *Server) initializeroutes() {
	server.Router.HandleFunc("/", controllers.Home).Methods("GET")
}
