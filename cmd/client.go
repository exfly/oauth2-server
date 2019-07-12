package cmd

import (
	"github.com/RichardKnop/go-oauth2-server/services"
)

// ClientManager management the client
func ClientManager(configBackend string) error {
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
	_, err = services.OauthService.CreateClient("app_1", "app_1_secret", "http://localhost:8081")
	_, err = services.OauthService.CreateUser("superuser", "superuser@test", "password")
	_, err = services.OauthService.CreateUser("user", "user@test", "password")
	return nil
}
