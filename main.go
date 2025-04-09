package main

import (
	"log"
	"microservicio/config"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ No se pudo cargar el archivo .env (modo producción)")
	}

	config.ConnectDB()
}
