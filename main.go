package main

import (
	"compressor/routes"
	"fmt"
	"log"
	"net/http"
)

//var port = os.Getenv("PORT")
var mux = http.NewServeMux()
var handler = routes.NewArchiveRouter()

func main() {

	mux.HandleFunc("/", routes.Index)
	mux.Handle("/archive", handler)
	fmt.Println("running on")
	log.Fatal(http.ListenAndServe(":8080", mux))

}
