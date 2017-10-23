package main

import (
	"compressor/routes"
	"net/http"
	"os"
)

var port = os.Getenv("PORT")

func main() {
	handler := routes.NewArchiveRouter()
	mux.Handle("/archive", handler)
	http.ListenAndServe(port, mux)

}

var mux = http.NewServeMux()
