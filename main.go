package main

import (
	"log"
	"net/http"

	"github.com/cenkayla/votingservice/service"
)

func main() {
	app := service.Service{}
	if err := app.Open(); err != nil {
		log.Fatal("open failed", err)
	}

	srv := &http.Server{
		Addr:    ":8080",
		Handler: app.ConfigureRouter(),
	}

	srv.ListenAndServe()
}
