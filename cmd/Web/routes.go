package main

import (
	"net/http"
	"github.com/virattGoogle/bookings/pkg/config"
	handeler "github.com/virattGoogle/bookings/pkg/handelrs"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func route(app *config.Appconfig) http.Handler{

	//mux := pat.New()

	mux:= chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(middleware.Logger)
	mux.Use(Nosurf)
	mux.Use(sessionLoad)
    mux.Get("/",handeler.Home)

	mux.Get("/About",handeler.About)
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*",http.StripPrefix("./static",fileServer))


	return mux
}