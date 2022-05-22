package app

import (
	"flag"
	"log"
	"os"

	"github.com/Jehanv60/gotoko/app/controllers"
	"github.com/joho/godotenv"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback

}

func Run() {
	var server = controllers.Server{}
	var appconfig = controllers.Appconfig{
		getEnv("APP_NAME", "Gotokoapp"),
		getEnv("APP_ENV", "Development"),
		getEnv("APP_PORT", "9000"),
	}
	dbconfig := controllers.DBconfig{
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

	flag.Parse()
	arg := flag.Arg(0)
	if arg != "" {
		server.Initcommand(appconfig, dbconfig)
	} else {
		server.Initialize(appconfig, dbconfig)
		server.Run(":" + appconfig.Appport)
	}
}
