package main

import (
	"fmt"
	"log"
	"net/http"

	"gitlab.com/cotrutatiberiu/project-go/src/app"
)

func main() {

	s := app.CreateServer()
	
	fmt.Println("Starting server on port: 5111")
	

	log.Fatal(http.ListenAndServe("127.0.0.1:5111", s.Router))
}