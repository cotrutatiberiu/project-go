package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/cotrutatiberiu/project-go/src/app"
)

func main() {
	os.Setenv("DATABASE_URL", "host=localhost port=5433 user=dbuser password=testdb dbname=opiniatadb sslmode=disable")

	s := app.CreateServer()

	fmt.Println("Starting server on port: 5111")

	log.Fatal(http.ListenAndServe("127.0.0.1:5111", s.Router))
}
