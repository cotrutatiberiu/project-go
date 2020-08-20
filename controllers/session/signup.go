package session

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/cotrutatiberiu/project-go/models"
)

type signupRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	UserName  string `json:"userName"`
	Email     string `json:"email"`
	Language  string `json:"language"`
	Password  string `json:"password"`
	Created   int    `json:"created"`
}

type signupResponse struct {
	Success bool `json:"success"`
}

func (c *Controller) Signup(w http.ResponseWriter, r *http.Request) {
	err := func(w http.ResponseWriter, r *http.Request) error {

		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return err
		}
		defer r.Body.Close()

		// Unmarshal
		acct := signupRequest{}
		err = json.Unmarshal(b, &acct)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return err
		}
		account := models.Account{
			AccountID: 0,
			FirstName: acct.FirstName,
			LastName:  acct.LastName,
			UserName:  acct.UserName,
			Email:     acct.Email,
			Language:  acct.Language,
		}
		err = c.acct.Create(r.Context(), &account)
		if err != nil {

			return err
		}

		resp := signupResponse{
			Success: true,
		}
		return json.NewEncoder(w).Encode(resp)
	}(w, r)
	if err != nil {
		log.Printf("session.Signup %s", err)
	}
}
