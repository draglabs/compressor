package main

import (
	"compressor/routes"
	"fmt"
	"log"
	"net/http"
)

//var port = os.Getenv("PORT")

func main() {
	handler := routes.NewArchiveRouter()
	mux.HandleFunc("/", routes.Index)
	mux.Handle("/archive", handler)
	log.Fatal(http.ListenAndServe(":8081", mux))
	fmt.Println("running on")
}

var mux = http.NewServeMux()
