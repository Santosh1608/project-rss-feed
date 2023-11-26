package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
}

func main() {
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*"},
		AllowedMethods: []string{"POST"},
	}))
	r.Use(middleware.Logger)

	PORT := os.Getenv("PORT")
	log.Println("⚡️ server started on port: " + PORT)
	err := http.ListenAndServe(":"+PORT, r)
	if err != nil {
		log.Fatal(err)
	}
}
