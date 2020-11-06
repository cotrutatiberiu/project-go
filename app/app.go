package app

import (
	"fmt"
	"net/http"

	"github.com/matryer/way"

	"github.com/cotrutatiberiu/project-go/repository/database"

	signupController "github.com/cotrutatiberiu/project-go/controllers/signup"
	accountRepository "github.com/cotrutatiberiu/project-go/repository/database/account"
	accountService "github.com/cotrutatiberiu/project-go/service/account"

	"github.com/cotrutatiberiu/project-go/utils"
)

type Server struct {
	Router *way.Router
	db     *database.Database
	// email
}

func CreateServer() *Server {

	utils.New()

	server := Server{
		Router: way.NewRouter(),
		db:     &database.Database{},
	}

	server.createRoutes()
	server.db.Setup()

	server.Router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "This is not the page you are looking for")
	})

	return &server
}

func (server *Server) createRoutes() {

	signupHandler := signupController.New(accountService.New(accountRepository.AccountRepository(), server.db))

	server.Router.HandleFunc(http.MethodPost, "/api/signup", signupHandler.HandlePost)
	// s.Router.HandleFunc("GET", "/api/signup/:id", signup.HandleSignup2())
	// s.Router.HandleFunc("GET", "/api/login", login.HandleSignup2())
}
