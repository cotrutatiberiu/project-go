package service

import "github.com/cotrutatiberiu/project-go/models"

type Validation interface {
	ValidateSignup(signupPayload models.SignupPayload) (bool, error)
}


