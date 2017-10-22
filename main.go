package main

import "net/http"
import "fmt"

func main() {
	mux.HandleFunc("/", handleArchive)
	http.ListenAndServe(":8080", mux)
}

var mux = http.NewServeMux()

func handleArchive(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "should handle")
}
func parseParams(r *http.Request) {

}
