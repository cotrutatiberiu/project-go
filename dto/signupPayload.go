package dto

import (
	"time"

	"github.com/go-playground/validator/v10"

	"github.com/cotrutatiberiu/project-go/models"
)

type SignupPayload struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	UserName  string `json:"userName"`
	Email     string `json:"email"`
	Language  string `json:"language"`
	Password  string `json:"password" validate:"min=8"`
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

// Direct validation
func (payload SignupPayload) Validate() error {
	return validator.New().Struct(payload)
}
