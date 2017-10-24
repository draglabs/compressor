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

// i need to put this middle ware before the interface method ServeHTTP
func middleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL)
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func (ar ArchiveRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		w.Header().Set("Content-Type", "application/json")
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
		// w.WriteHeader(http.StatusBadRequest)
		// json.NewEncoder(w).Encode(responseMessage{"cant parse params"})
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

// Index func, main index handler
func Index(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(responseMessage{"Compressor homepage"})
}
