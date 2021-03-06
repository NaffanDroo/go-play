package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

const (
	timeFormat = "15:04:05"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))

	// Routes
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		render.HTML(w, r, paragraph("Welcome, why not try "+bold("/v1/time")))
	})

	r.Route("/v1", func(r chi.Router) {
		r.Get("/time", getTime)
		r.Get("/hostname", getHostName)
	})

	fmt.Println("Listening...")
	http.ListenAndServe(":3000", r)
}

func getTime(w http.ResponseWriter, r *http.Request) {
	var now string = time.Now().Format(timeFormat)
	render.JSON(w, r, map[string]string{"time": now})
}

func getHostName(w http.ResponseWriter, r *http.Request) {
	name, _ := os.Hostname()
	render.JSON(w, r, map[string]string{"hostname": name})
}
