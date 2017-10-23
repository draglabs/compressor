package main

import "net/http"
import "compressor/routes"

func main() {
	handler := routes.NewArchiveRouter()
	mux.Handle("/archive", handler)
	http.ListenAndServe(":8080", mux)
}

var mux = http.NewServeMux()

func parseParams(r *http.Request) {

}
