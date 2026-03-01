package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Chirag711/go-rest-api/config"
	"github.com/Chirag711/go-rest-api/routes"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env")
	}

	config.ConnectDB()

	r := mux.NewRouter()
	routes.RegisterUserRoutes(r)

	log.Println("Server running on port", os.Getenv("PORT"))
	http.ListenAndServe(":"+os.Getenv("PORT"), r)
}
