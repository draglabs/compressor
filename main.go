package main

import (
	"compressor/routes"
	"net/http"
)

func main() {
	handler := routes.NewArchiveRouter()
	mux.Handle("/archive", handler)
	http.ListenAndServe(":8080", mux)

}

var mux = http.NewServeMux()
