package controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Jehanv60/gotoko/app/models"
	"github.com/Jehanv60/gotoko/database/seeder"
	"github.com/gorilla/mux"
	"github.com/urfave/cli"
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
	// seeder.DBseed(server.DB)
}

func (server *Server) initializeDB(dbconfig DBconfig) {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", dbconfig.DBHost, dbconfig.DBuser, dbconfig.DBPassword, dbconfig.DBName, dbconfig.DBPort)
	server.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("failed conenct database")
	}

}

func (server *Server) dbMigrate() {
	for _, model := range models.RegisterModels() {
		err := server.DB.Debug().AutoMigrate(model.Model)
		if err != nil {
			log.Fatalln(err)
		}
	}
	fmt.Println("proses migarsi sukses")
}
func (server *Server) Initcommand(config Appconfig, dbconfig DBconfig) {
	server.initializeDB(dbconfig)
	cmdapp := cli.NewApp()
	cmdapp.Commands = []cli.Command{
		{
			Name: "db:migrate",
			Action: func(c *cli.Context) error {
				server.dbMigrate()
				return nil
			},
		},
		{
			Name: "db:seed",
			Action: func(c *cli.Context) error {
				err := seeder.DBseed(server.DB)
				if err != nil {
					log.Fatalln(err)
				}
				return nil
			},
		},
	}
	err := cmdapp.Run(os.Args)
	if err != nil {
		log.Fatalln(err)
	}
}

func (server *Server) Run(addr string) {
	fmt.Printf("listening to port %s", addr)
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
