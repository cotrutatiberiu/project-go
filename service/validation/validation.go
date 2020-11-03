package validation

import (
	"errors"

	"github.com/cotrutatiberiu/project-go/models"
	"github.com/cotrutatiberiu/project-go/service"
)

type validationService struct {
}

func New() service.Validation {
	return &validationService{}
}

func (service *validationService) ValidateSignup(signupPayload models.SignupPayload) (bool, error) {
	return true, errors.New("asd")
}
