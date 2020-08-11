package app

import (
	"fmt"
	"net/http"

	"github.com/cotrutatiberiu/project-go/controllers/account"
	"github.com/cotrutatiberiu/project-go/controllers/session"
	"github.com/cotrutatiberiu/project-go/db"
	"github.com/cotrutatiberiu/project-go/repository/database"
	"github.com/jackc/pgx/v4"
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

	// TODO: real db setup here
	var db *pgx.Conn

	acct := account.New(database.Account(db))
	// Account CRUD operations
	r.HandleFunc(http.MethodGet, "/api/account/", acct.ListAll)
	r.HandleFunc(http.MethodGet, "/api/account/:id", acct.List)
	r.HandleFunc(http.MethodPost, "/api/account/", acct.Create)
	r.HandleFunc(http.MethodPut, "/api/account/:id", acct.Update)
	r.HandleFunc(http.MethodDelete, "/api/account/:id", acct.Update)

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
