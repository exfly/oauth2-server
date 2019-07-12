package cmd

import (
	"log"
	"net/http"
	"time"

	"github.com/RichardKnop/go-oauth2-server/services"
	"github.com/gorilla/mux"
	"github.com/phyber/negroni-gzip/gzip"
	"github.com/urfave/negroni"
	graceful "gopkg.in/tylerb/graceful.v1"
)

// RunServer runs the app
func RunServer(configBackend string) error {
	cnf, db, err := initConfigDB(true, true, configBackend)
	if err != nil {
		return err
	}
	defer db.Close()

	// start the services
	if err := services.Init(cnf, db); err != nil {
		return err
	}
	defer services.Close()

	// Start a classic negroni app
	app := negroni.New()
	app.Use(negroni.NewRecovery())
	app.Use(negroni.NewLogger())
	app.Use(gzip.Gzip(gzip.DefaultCompression))
	app.Use(negroni.NewStatic(http.Dir("public")))

	// Create a router instance
	router := mux.NewRouter()

	// Add routes
	services.HealthService.RegisterRoutes(router, "/v1")
	services.OauthService.RegisterRoutes(router, "/v1/oauth")
	services.WebService.RegisterRoutes(router, "/web")

	// Set the router
	app.UseHandler(router)

	log.Println("runserver at :8080")
	// Run the server on port 8080, gracefully stop on SIGTERM signal
	graceful.Run(":8080", 5*time.Second, app)

	return nil
}
