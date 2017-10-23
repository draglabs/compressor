package routes

import (
	"compressor/controllers"
	"compressor/models"
	"encoding/json"
	"fmt"
	"net/http"
)

// ArchiveRouter struct
type ArchiveRouter struct {
}
type responseMessage struct {
	Message string `json:"message"`
}

// NewArchiveRouter func
func NewArchiveRouter() ArchiveRouter {
	return ArchiveRouter{}
}
func (ar ArchiveRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":

		handleArchive(w, r)
	default:
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte("method not implemented"))

		return
	}
}
func handleArchive(w http.ResponseWriter, r *http.Request) {
	jam, err := controllers.FetchJam(parseParams(r))

	if err == nil {
		er := json.NewEncoder(w).Encode(&jam)
		fmt.Println(er)
	}
}

func parseParams(r *http.Request) *models.ArchiveParam {

	userID := r.FormValue("user_id")
	jamID := r.FormValue("jam_id")
	if userID != "" && jamID != "" {
		return &models.ArchiveParam{UserID: userID, JamID: jamID}
	}
	return nil
}
