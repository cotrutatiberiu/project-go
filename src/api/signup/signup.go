package signup

import (
	"fmt"
	"net/http"

	"gitlab.com/cotrutatiberiu/project-go/src/db"
)

type signupPayload struct {
	firstName string
	lastName  string
	userName  string
	email     string
	language  string
	password  string
	created   int
	updated   int
}

func HandleSignup() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// result := db.Create("SELECT email FROM accounts WHERE account_id = 7")
		// values := signupPayload{"a", "a", "a", "a", "a", "a", 1596747588, 1596747588}
		result := db.Create(`INSERT INTO 
		accounts(first_name, last_name, user_name, email, language, password, created, updated)
		VALUES($1, $2, $3, $4, $5, $6, $7, $8);`)

		fmt.Println(result)
	}
}

// func HandleSignup2() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		// fmt.Println(r)
// 		id := way.Param(r.Context(), "id")
// 		fmt.Println(id)

// 		s := db.ExecuteQuery("SELECT email FROM accounts WHERE account_id = 7")
// 		fmt.Println(s)
// 	}
// }

func handleStuff(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "helo")
}
