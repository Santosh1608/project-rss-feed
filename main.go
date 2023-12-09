package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/santosh1608/project-rss/dynamo"
	"github.com/santosh1608/project-rss/handlers"
)

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
	dynamo.GetClient()
	// go startScraping(time.Second * 5)
}

func main() {

	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*"},
		AllowedMethods: []string{"POST"},
	}))
	r.Use(middleware.Logger)
	r.Use(middleware.Heartbeat("/heathz"))
	r.Mount("/api", routes())

	PORT := os.Getenv("PORT")
	log.Println("⚡️ server started on port: " + PORT)
	err := http.ListenAndServe(":"+PORT, r)
	if err != nil {
		log.Fatal(err)
	}
}

func routes() *chi.Mux {
	apiRouter := chi.NewRouter()

	// public routes
	apiRouter.Group(func(r chi.Router) {
		r.Post("/register", handlers.Register)
		r.Post("/login", handlers.Login)
	})

	// protected routes
	apiRouter.Group(func(r chi.Router) {
		// r.Use(jwtauth.Verifier(auth.TokenAuth))
		// r.Use(jwtauth.Authenticator(auth.TokenAuth))

		r.Post("/feed", handlers.CreateFeed)
		r.Post("/follow/{feedId}", handlers.FollowFeed)
		r.Get("/posts/{feedId}", handlers.FetchPosts)
		r.Get("/posts/{postId}/feed/{feedId}", handlers.FetchPostByFeedId)
	})
	return apiRouter
}
