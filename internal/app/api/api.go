package api

import (
	"github.com/narekps/RestApi/internal/app/models"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Api struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	data   *models.Data
}

func New(config *Config) *Api {
	return &Api{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// Start starting http server
func (api *Api) Start() error {
	if err := api.configureLogger(); err != nil {
		return err
	}

	api.configureRouter()

	api.logger.Info("Starting api server at port ", api.config.BindAddr)

	return http.ListenAndServe(api.config.BindAddr, api.router)
}
