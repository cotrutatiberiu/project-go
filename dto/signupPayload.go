package dto

import (
	"time"

	"github.com/cotrutatiberiu/project-go/models"
)

type SignupPayload struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	UserName  string `json:"userName"`
	Email     string `json:"email"`
	Language  string `json:"language"`
	Password  string `json:"password"`
}

func (payload SignupPayload) ToModel() *models.Account {
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
