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
// represents the single router we have
func NewArchiveRouter() ArchiveRouter {
	return ArchiveRouter{}
}

// i need to put this middleware before the interface method ServeHTTP
func middleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL)
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

// ServeHTTP func, is the interface conformance
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

// handleArchive func, is the handler for the
// archive route, it fetches the jam by a given id
// and sends it back to the user
func handleArchive(w http.ResponseWriter, r *http.Request) {
	jam, err := controllers.FetchJam(parseParams(r))
	if err == nil {
		er := json.NewEncoder(w).Encode(&jam)
		if er != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(responseMessage{"something when wrong: " + er.Error()})
		}
	}
}

// parseParams func, parses the params of the incoming
// request and checks for simple validation
func parseParams(r *http.Request) *models.ArchiveParam {
	var para models.ArchiveParam
	err := json.NewDecoder(r.Body).Decode(&para)
	defer r.Body.Close()
	if err == nil {
		fmt.Println(para.JamID)
		return &para
	}

	return nil
}

// Index func, main index handler
// handles the main entry point
// by default it sends a message in json
// to tell the user she/he as riched the index
func Index(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(responseMessage{"Compressor homepage"})
}
