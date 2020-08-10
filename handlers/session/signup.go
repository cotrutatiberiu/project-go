package session

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/cotrutatiberiu/project-go/db"
)

type signupPayload struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	UserName  string `json:"userName"`
	Email     string `json:"email"`
	Language  string `json:"language"`
	Password  string `json:"password"`
	Created   int    `json:"created"`
}

func HandleSignup() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		b, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		// Unmarshal
		var payload signupPayload
		err = json.Unmarshal(b, &payload)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		fmt.Println(payload)
		// result := db.Create("SELECT email FROM accounts WHERE account_id = 7")
		// values := signupPayload{"a", "a", "a", "a", "a", "a", 1596747588, 1596747588}

		pool := db.CreateConnection()
		result := db.Create(pool, `INSERT INTO
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
