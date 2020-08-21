package session

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// loginRequest expected json payload
type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// loginResponse a dummy response
type loginResponse struct {
	Success bool `json:"success"`
}

func (c *Controller) Login(w http.ResponseWriter, r *http.Request) {
	err := func(w http.ResponseWriter, r *http.Request) error {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return err
		}
		defer r.Body.Close()

		creds := loginRequest{}
		err = json.Unmarshal(body, &creds)
		if err != nil {
			return err
		}
		token, err := c.svc.Login(r.Context(), creds.Username, creds.Password)
		if err != nil {
			// TODO improve this
			return json.NewEncoder(w).Encode(loginResponse{})
		}

		// TODO set cookie
		log.Printf("new token  %s", token)

		resp := loginResponse{
			Success: true,
		}
		return json.NewEncoder(w).Encode(resp)
	}(w, r)
	if err != nil {
		log.Printf("session.Login %s", err)
	}
}
