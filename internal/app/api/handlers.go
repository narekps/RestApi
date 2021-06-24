package api

import (
	"encoding/json"
	"github.com/narekps/RestApi/internal/app/models"
	"net/http"
)

type ResultSet struct {
	Success bool
	Code    int
	Message string
}

func initHeaders(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
}

func writeResultSet(writer http.ResponseWriter, success bool, code int, message string) {
	resultSet := ResultSet{
		Success: success,
		Code:    code,
		Message: message,
	}
	writer.WriteHeader(resultSet.Code)
	json.NewEncoder(writer).Encode(resultSet)
}

// GET /ping
func (api *Api) pingHandler(writer http.ResponseWriter, request *http.Request) {
	api.logger.Info(request.RequestURI)

	initHeaders(writer)
	json.NewEncoder(writer).Encode("pong")
}

// POST /grab
func (api *Api) grabHandler(writer http.ResponseWriter, request *http.Request) {
	api.logger.Info(request.RequestURI)

	initHeaders(writer)

	var data models.Data
	err := json.NewDecoder(request.Body).Decode(&data)
	if err != nil {
		api.logger.Warning("Invalid json received from client")
		writeResultSet(writer, false, http.StatusBadRequest, "Provided json is invalid")
		return
	}

	api.data = &data

	writeResultSet(writer, true, http.StatusCreated, "Data accepted")
}

// GET /solve
func (api *Api) solveHandler(writer http.ResponseWriter, request *http.Request) {
	api.logger.Info(request.RequestURI)

	initHeaders(writer)

	if api.data == nil {
		api.logger.Warning("Data not defined")
		writeResultSet(writer, false, http.StatusBadRequest, "Data not defined")
		return
	}

	api.data.Calculate()

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(api.data)
}
