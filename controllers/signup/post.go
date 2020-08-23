package signup

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/cotrutatiberiu/project-go/models"
)

type signupPayload struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	UserName  string `json:"userName"`
	Email     string `json:"email"`
	Language  string `json:"language"`
	Password  string `json:"password"`
	Created   int64  `json:"created"`
}

func (payload signupPayload) ToModel() *models.Account {
	return &models.Account{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		UserName:  payload.UserName,
		Email:     payload.Email,
		Language:  payload.Language,
		Password:  payload.Password,
		Created:   time.Now(),
		Updated:   time.Now(),
	}
}

func (controller *Controller) HandlePost(w http.ResponseWriter, r *http.Request) {
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

	account := payload.ToModel()

	//TODO: see whatsup with request.context
	controller.accountService.Create(context.Background(), account)
}
