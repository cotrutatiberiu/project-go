package dto

import (
	"time"

	"github.com/go-playground/validator/v10"

	"github.com/cotrutatiberiu/project-go/models"
	"github.com/cotrutatiberiu/project-go/utils"
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

func (payload SignupPayload) Validate() []string {
	err := validator.New().Struct(payload)

	var errTags []string

	for _, err := range err.(validator.ValidationErrors) {
		utils.Instance.Logger.Printf("%s %s %s %s - %s\n", err.Namespace(), err.Type(), err.Tag(), err.Param(), err.Value())

		errTags = append(errTags, err.Tag())
	}

	return errTags
}
