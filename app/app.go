package app

import (
	"fmt"
	"net/http"

	"github.com/cotrutatiberiu/project-go/controllers/account"
	"github.com/cotrutatiberiu/project-go/controllers/session"
	"github.com/cotrutatiberiu/project-go/repository/database"
	account2 "github.com/cotrutatiberiu/project-go/service/account"
	session2 "github.com/cotrutatiberiu/project-go/service/session"
	"github.com/jackc/pgx/v4"
	"github.com/matryer/way"
)

type Server struct {
	Router *way.Router
	// email
}

func CreateServer() *Server {
	r := way.NewRouter()

	server := Server{
		Router: r,
	}

	// TODO: real db setup here
	var db *pgx.Conn

	acctRepo := database.Account(db)
	acctSvc := account2.New(acctRepo)

	acct := account.New(acctSvc)
	// Account CRUD operations
	r.HandleFunc(http.MethodGet, "/api/account/", acct.ListAll)
	r.HandleFunc(http.MethodGet, "/api/account/:id", acct.List)
	r.HandleFunc(http.MethodPost, "/api/account/", acct.Create)
	r.HandleFunc(http.MethodPut, "/api/account/:id", acct.Update)
	r.HandleFunc(http.MethodDelete, "/api/account/:id", acct.Update)

	sess := session.New(session2.New(acctRepo), acctSvc)
	// Session operations
	r.HandleFunc(http.MethodPost, "/api/signup", sess.Signup)
	r.HandleFunc(http.MethodPost, "/api/login", sess.Login)

	server.Router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "This is not the page you are looking for")
	})

	return &server
}
