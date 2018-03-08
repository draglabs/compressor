package main

import (
	"compressor/routes"
	"fmt"
	"log"
	"net/http"
	"os"
)

var port = os.Getenv("PORT")
var mux = http.NewServeMux()
var handler = routes.NewArchiveRouter()

func main() {

	mux.HandleFunc("/", routes.Index)
	mux.Handle("/archive", handler)
	fmt.Println("running on")
	if prod := os.Getenv("PROD"); prod == "" {
		port = ":8081"
		fmt.Println(port)
	}
	log.Fatal(http.ListenAndServe(port, mux))

}
