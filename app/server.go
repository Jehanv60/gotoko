package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"gorm.io/gorm"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

type Appconfig struct {
	Appname string
	Appenv  string
	Appport string
}

func (server *Server) Initalize(Appconfig Appconfig) {
	fmt.Println("welcome to " + Appconfig.Appname)
	server.Router = mux.NewRouter()
	server.initializeroutes()
}

func (server *Server) Run(addr string) {
	fmt.Printf("listening to port %s", addr)
	log.Fatal(http.ListenAndServe(addr, server.Router))
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback

}

func Run() {
	var server = Server{}
	var appconfig = Appconfig{
		getEnv("APP_Name", "Gotokoapp"),
		getEnv("APP_ENV", "Development"),
		getEnv("APP_PORT", "9000")}

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loading env file")
	}
	server.Initalize(appconfig)
	server.Run(":" + appconfig.Appport)
}
