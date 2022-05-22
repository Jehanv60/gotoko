package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Jehanv60/gotoko/app/models"
	"github.com/Jehanv60/gotoko/database/seeder"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"gorm.io/driver/postgres"
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

type DBconfig struct {
	DBHost     string
	DBuser     string
	DBPassword string
	DBName     string
	DBPort     string
}

func (server *Server) Initialize(Appconfig Appconfig, dbconfig DBconfig) {
	fmt.Println("welcome to " + Appconfig.Appname)
	server.initializeDB(dbconfig)
	server.initializeroutes()
	seeder.DBseed(server.DB)
}

func (server *Server) initializeDB(dbconfig DBconfig) {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", dbconfig.DBHost, dbconfig.DBuser, dbconfig.DBPassword, dbconfig.DBName, dbconfig.DBPort)
	server.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("failed conenct database")
	}

	for _, model := range models.RegisterModels() {
		err := server.DB.Debug().AutoMigrate(model.Model)
		if err != nil {
			log.Fatalln(err)
		}
	}
	fmt.Println("proses migarsi sukses")
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
		getEnv("APP_NAME", "Gotokoapp"),
		getEnv("APP_ENV", "Development"),
		getEnv("APP_PORT", "9000"),
	}
	dbconfig := DBconfig{
		getEnv("DB_HOST", "localhost"),
		getEnv("DB_USER", "han"),
		getEnv("DB_PASSWORD", "solo"),
		getEnv("DB_NAME", "gotoko"),
		getEnv("APP_PORT", "5432"),
	}

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loading env file")
	}
	server.Initialize(appconfig, dbconfig)
	server.Run(":" + appconfig.Appport)
}
