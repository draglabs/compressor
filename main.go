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
	mux.HandleFunc("/", routes.Index)
	mux.Handle("/archive", handler)
	log.Fatal(http.ListenAndServe(port, mux))

}

var mux = http.NewServeMux()
