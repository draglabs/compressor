package main

import (
	"compressor/routes"
	"log"
	"net/http"
	"os"
)

var port = os.Getenv("PORT")

func main() {
	handler := routes.NewArchiveRouter()
	mux.Handle("/archive", handler)
	log.Fatal(http.ListenAndServe(port, mux))

}

var mux = http.NewServeMux()
