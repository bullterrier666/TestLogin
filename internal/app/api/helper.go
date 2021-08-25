package api

import (
	"net/http"

	"github.com/bullterrier666/TestLogin/internal/app/middleware"
	_ "github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

var (
	prefix string = "/api"
)

func (a *API) configureLoggerField() error {
	log_level, err := logrus.ParseLevel(a.config.LoggerLevel)
	if err != nil {
		return err
	}
	a.logger.SetLevel(log_level)
	return nil
}

func (a *API) configureRouterField() {
	//For authorization
	a.router.HandleFunc(prefix+"/user/auth", a.PostToAuth).Methods("POST")
	a.router.Handle(prefix+"/test", middleware.JwtMiddleware.Handler(
		http.HandlerFunc(a.TestAuth),
	)).Methods("GET")

}
