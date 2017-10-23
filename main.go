package main

import (
	"compressor/mailer"
	"compressor/routes"
	"net/http"
)

func main() {
	mailer.SendMail()
	handler := routes.NewArchiveRouter()
	mux.Handle("/archive", handler)
	http.ListenAndServe(":8080", mux)

}

var mux = http.NewServeMux()
