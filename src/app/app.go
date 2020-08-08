package app

import (
	"fmt"
	"net/http"

	"github.com/matryer/way"
	"gitlab.com/cotrutatiberiu/project-go/src/api/signup"
	// "gitlab.com/cotrutatiberiu/project-go/src/api/login"
)

type Server struct {
	Router *way.Router
	// db
	// email
}

func CreateServer() *Server {
	s := &Server{way.NewRouter()}

	s.createRoutes()

	s.Router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "This is not the page you are looking for")
	})

	return s
}

func (s *Server) createRoutes() {
	s.Router.HandleFunc("GET", "/api/signup", signup.HandleSignup())
	// s.Router.HandleFunc("GET", "/api/signup/:id", signup.HandleSignup2())
	// s.Router.HandleFunc("GET", "/api/login", login.HandleSignup2())
}
