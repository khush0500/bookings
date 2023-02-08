package main

import (
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/khush0500/bookings/pkg/config"
	"github.com/khush0500/bookings/pkg/handlers"
	"github.com/khush0500/bookings/pkg/renders"
)

var app config.AppConfig

var session *scs.SessionManager

func main() {

	//cahnge this to true when in Production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false

	app.Session = session

	tc, err := renders.CreateTemplateCache()
	// fmt.Println(tc)
	if err != nil {
		log.Fatal("Cannot create Template cache")
	}

	app.TemplateCache = tc
	app.UseCache = app.InProduction

	repo := handlers.NewRepo(&app)
	handlers.NewHanlers(repo)

	renders.NewTemplates(&app)
	srv := &http.Server{
		Addr:    ":8080",
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
