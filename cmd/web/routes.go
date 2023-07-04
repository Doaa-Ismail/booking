package main

import (
	"net/http"

	"github.com/Doaa-Ismail/go_course/booking/pkg/config"
	"github.com/Doaa-Ismail/go_course/booking/pkg/handlers"
	"github.com/bmizerany/pat"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func RoutesPat(app *config.AppConfig) http.Handler {
	mux := pat.New()
	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
	mux.Get("/check", http.HandlerFunc(handlers.Repo.Divide))
	mux.Get("/hello", http.HandlerFunc(handlers.Repo.Hello))
	return mux
}

func RoutesChi(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	//mux.Use(writeToConsole)
	mux.Use(NoSurf)

	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/check", handlers.Repo.Divide)
	mux.Get("/hello", handlers.Repo.Hello)
	return mux
}
