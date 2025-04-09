package main

import (
	"log"
	"microservicio/config"
	"microservicio/routes"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ .env no encontrado, se esperan variables del sistema")
	}

	config.ConnectDB()
	router := routes.SetupRoutes()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("🚀 Servidor corriendo en http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
