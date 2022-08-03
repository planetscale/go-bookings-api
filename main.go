package main

import (
	"net/http"
	"os"

	"bookings-api/routes"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("welcome"))
	})

	r.Route("/hotels", routes.Hotels)

	http.ListenAndServe(os.Getenv("LISTEN"), r)
}
