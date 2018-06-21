package main

import (
	"github.com/draglabs/compressor/routes"
	"fmt"
	"log"
	"net/http"
	"os"

	"context"
	"github.com/MindsightCo/go-mindsight-collector"
)

var port = os.Getenv("PORT")
var mux = http.NewServeMux()
var handler = routes.NewArchiveRouter()

func main() {

	// functions to start mindsight
	ctx := context.Background()
	collector.StartMindsightCollector(ctx,
		collector.OptionAgentURL("http://localhost:8000/samples/"),
		collector.OptionWatchPackage("github.com/draglabs/compressor/"))

	// functions to start the server
	mux.HandleFunc("/", routes.Index)
	mux.Handle("/archive", handler)
	fmt.Println("running on")
	if prod := os.Getenv("PROD"); prod == "" {
		port = ":8081"
	}
	log.Fatal(http.ListenAndServe(port, mux))

}
