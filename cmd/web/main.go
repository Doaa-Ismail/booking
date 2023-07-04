package main

// go run cmd/web/main.go cmd/web/routes.go cmd/web/middleware.go

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Doaa-Ismail/go_course/booking/pkg/config"
	"github.com/Doaa-Ismail/go_course/booking/pkg/handlers"
	"github.com/Doaa-Ismail/go_course/booking/pkg/render"

	"github.com/alexedwards/scs/v2"
)

const PortNumber = ":8080"

var app config.AppConfig

var session *scs.SessionManager

func main() {
	fmt.Println("Web is running with Porting Number : ", PortNumber)
	//false when development,, true if In Production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateRenderTemplatesTest()

	if err != nil {
		log.Fatal("Cannot create template cache")
	}
	app.TemplateCache = tc

	app.UseCache = false

	repo := handlers.NewRepository(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)
	// http.HandleFunc("/check", handlers.Repo.Divide)
	// http.HandleFunc("/hello", handlers.Repo.Hello)
	// _ = http.ListenAndServe(PortNumber, nil)

	// srv := &http.Server{
	// 	Addr:    PortNumber,
	// 	Handler: RoutesPat(&app),
	// }
	// err = srv.ListenAndServe()
	// log.Fatal(err)

	srv := &http.Server{
		Addr:    PortNumber,
		Handler: RoutesChi(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)
}
