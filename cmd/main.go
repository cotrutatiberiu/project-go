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

	server := app.CreateServer()

	httpServer := &http.Server{
		Addr:         defaultHTTPAddr,
		Handler:      server.Router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	host, port, err := net.SplitHostPort(httpServer.Addr)
	if err != nil {
		log.Fatalf("invalid http listen address %s", err)
	}

	log.Printf("Starting server %s on port: %s", host, port)

	err = httpServer.ListenAndServe()
	if err != nil {
		log.Fatalf("failed to start http service %s", err)
	}
}
