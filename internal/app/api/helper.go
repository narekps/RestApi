package api

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

func (api *Api) configureLogger() error {
	logLevel, err := logrus.ParseLevel(api.config.LoggerLevel)
	if err != nil {
		return nil
	}

	api.logger.SetLevel(logLevel)

	return nil
}

func (api *Api) configureRouter() {
	api.router.HandleFunc("/ping", api.pingHandler).Methods(http.MethodGet)
	api.router.HandleFunc("/grab", api.grabHandler).Methods(http.MethodPost)
	api.router.HandleFunc("/solve", api.solveHandler).Methods(http.MethodGet)
}
