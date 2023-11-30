package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"github.com/virattGoogle/bookings/pkg/config"
	handeler "github.com/virattGoogle/bookings/pkg/handelrs"
	"github.com/virattGoogle/bookings/pkg/render"

	"github.com/alexedwards/scs/v2"
)

const port = ":8080"

var app config.Appconfig

var session *scs.SessionManager

func main() {


app.InProduction = false
	session = scs.New()
	session.Lifetime = 24* time.Hour
	session.Cookie.Persist = true
    session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session


	tc,err := render.Createttemplatecache()

	if err != nil{
		log.Println(err)
	}

	repo := handeler.NewRepo(&app)

	handeler.NewHandler(repo)

	app.TemplateCache = tc

	render.Newtemplate(&app)

	//http.HandleFunc("/", handeler.Home)
	//http.HandleFunc("/About", handeler.About)

	fmt.Println("server listenung on %s", port)

	//_ = http.ListenAndServe(port, nil)

	serve := &http.Server{
		Addr: port,
		Handler: route(&app),

	}

	err = serve.ListenAndServe()

	log.Println(err)

}
