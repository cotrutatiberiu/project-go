package app

import (
	"fmt"
	"net/http"

	"github.com/cotrutatiberiu/project-go/db"
	"github.com/cotrutatiberiu/project-go/handlers/account"
	"github.com/cotrutatiberiu/project-go/handlers/session"
	"github.com/matryer/way"
)

type Server struct {
	Router *way.Router
	db     *db.Database
	// email
}

func CreateServer() *Server {
	r := way.NewRouter()

	server := Server{
		Router: r,
		db:     &db.Database{},
	}

	server.db.Setup()

	// Account CRUD operations
	r.HandleFunc(http.MethodGet, "/api/account/", account.List)
	r.HandleFunc(http.MethodGet, "/api/account/:id", account.Get)
	r.HandleFunc(http.MethodPost, "/api/account/", account.Create)
	r.HandleFunc(http.MethodPut, "/api/account/:id", account.Update)
	r.HandleFunc(http.MethodDelete, "/api/account/:id", account.Update)

	// Session operations
	r.HandleFunc("POST", "/api/signup", session.HandleLogin())
	r.HandleFunc("GET", "/api/signup/:id", session.HandleSignup2())
	r.HandleFunc("GET", "/api/login", session.HandleSignup2())

	server.Router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "This is not the page you are looking for")
	})

	return &server
}
