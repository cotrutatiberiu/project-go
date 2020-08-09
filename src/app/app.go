package app

import (
	"fmt"
	"net/http"

	"github.com/cotrutatiberiu/project-go/src/api/signup"
	"github.com/cotrutatiberiu/project-go/src/db"
	"github.com/matryer/way"
	// "github.com/cotrutatiberiu/project-go/src/api/login"
)

type Server struct {
	Router *way.Router
	db     *db.Database
	// email
}

func CreateServer() *Server {
	server := &Server{way.NewRouter(), &db.Database{}}

	server.createRoutes()
	server.db.Setup()

	server.Router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "This is not the page you are looking for")
	})

	return server
}

func (s *Server) createRoutes() {
	s.Router.HandleFunc("POST", "/api/signup", signup.HandleSignup())
	// s.Router.HandleFunc("GET", "/api/signup/:id", signup.HandleSignup2())
	// s.Router.HandleFunc("GET", "/api/login", login.HandleSignup2())
}
