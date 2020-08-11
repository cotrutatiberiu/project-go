package session

import (
	"fmt"
	"net/http"

	//"github.com/cotrutatiberiu/project-go/db"
	"github.com/matryer/way"
)

func HandleLogin() http.HandlerFunc {
	return handleStuff
}

func HandleSignup2() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := way.Param(r.Context(), "id")
		fmt.Println(id)

		//s := db.ExecuteQuery("SELECT email FROM accounts WHERE account_id = 7")
		//fmt.Println(s)

	}
}

func handleStuff(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "helo")
}
