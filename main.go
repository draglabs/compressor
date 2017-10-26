package main

import (
	"compressor/routes"
	"fmt"
	"log"
	"net/http"
)

var port = os.Getenv("PORT")

func main() {
	handler := routes.NewArchiveRouter()
	mux.HandleFunc("/", routes.Index)
	mux.Handle("/archive", handler)
	fmt.Println("running on")
	log.Fatal(http.ListenAndServe(port, mux))

}

var mux = http.NewServeMux()
