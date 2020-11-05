package service

import "github.com/cotrutatiberiu/project-go/dto"

type Validation interface {
	ValidateSignup(signupPayload dto.SignupPayload) (bool, error)
}
