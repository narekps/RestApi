package api

import "net/http"

// GET /ping
func (api *Api) pingHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("pong"))
	if err != nil {
		api.logger.Error(err)
	}
}
