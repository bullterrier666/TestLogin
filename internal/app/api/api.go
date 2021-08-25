package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type API struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	//storage *storage.Storage
}

func init() {}

func New(config *Config) *API {
	return &API{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (api *API) Start() error {
	//Trying to configure logger
	if err := api.configureLoggerField(); err != nil {
		return err
	}
	api.logger.Info("Starting api server at port:", api.config.BindAddr)
	api.configureRouterField()
	//if err := api.configureStorageField(); err != nil {
	//		return err
	//	}
	return http.ListenAndServe(api.config.BindAddr, api.router)
}
