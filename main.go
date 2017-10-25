package main

import (
	"compressor/routes"
	"log"
	"net/http"
)

//var port = os.Getenv("PORT")

func main() {
	handler := routes.NewArchiveRouter()
	mux.HandleFunc("/", routes.Index)
	mux.Handle("/archive", handler)
	log.Fatal(http.ListenAndServe(":8080", mux))
}

var mux = http.NewServeMux()
