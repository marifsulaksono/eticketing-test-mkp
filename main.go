package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/marifsulaksono/eticketing-test-mkp/config"
)

const PORT = "SERVER_PORT"

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading files .env : %v", err)
	}

	db := config.InitDatabase()
	r := InitRoutes(db)

	log.Printf("Starting server at localhost:%s .....", os.Getenv(PORT))
	err := http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv(PORT)), r)
	if err != nil {
		log.Fatalf("Error starting server : %v", err)
	}
}
