package main

import (
	
	"net/http"

	"github.com/justinas/nosurf"
)

func Nosurf(next http.Handler) http.Handler{

csrfHandeler := nosurf.New(next)

csrfHandeler.SetBaseCookie(http.Cookie{
	HttpOnly: true,
	Path: "/",
	Secure: app.InProduction,
	SameSite: http.SameSiteLaxMode,
})
return csrfHandeler
}

func sessionLoad(next http.Handler) http.Handler{
	return session.LoadAndSave(next)
}

