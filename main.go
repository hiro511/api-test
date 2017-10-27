package main

import (
	"log"
	"net/http"

	"github.com/hiro511/api-test/api"
)

const port = ":8080"

func main() {
	mux := http.NewServeMux()
	api.Register(mux)
	srv := http.Server{
		Addr:    port,
		Handler: mux,
	}
	srv.SetKeepAlivesEnabled(false)

	log.Printf("http server will run on %s", port)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
