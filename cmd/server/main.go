package main

import (
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/cotrutatiberiu/project-go/app"
)

const (
	defaultHTTPAddr = "127.0.0.1:5111"
)

func main() {
	os.Setenv("DATABASE_URL", "host=localhost port=5433 user=dbuser password=testdb dbname=opiniatadb sslmode=disable")

	s := app.CreateServer()

	srv := &http.Server{
		Addr:         defaultHTTPAddr,
		Handler:      s.Router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	_, port, err := net.SplitHostPort(srv.Addr)
	if err != nil {
		log.Fatalf("invalid http listen address %s", err)
	}

	log.Printf("starting server on port: %d", port)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatalf("failed to start http service %s", err)
	}

}
