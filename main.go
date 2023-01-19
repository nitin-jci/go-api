package main

import (
	"net/http"

	"github.com/joho/godotenv"
	"github.com/nitin-jci/go-api/controllers"
	"github.com/nitin-jci/go-api/models"
)

func main() {
	godotenv.Load()

	handler := controllers.New()

	server := &http.Server{
		Addr:    "0.0.0.0:8008",
		Handler: handler,
	}

	models.ConnectDatabase()

	server.ListenAndServe()
}
