package validation

import (
	"errors"
	"log"

	"github.com/cotrutatiberiu/project-go/models"
	"github.com/cotrutatiberiu/project-go/service"
	"github.com/go-playground/validator/v10"
)

type validationService struct {
}

func New() service.Validation {
	return &validationService{}
}

func (service *validationService) ValidateSignup(signupPayload models.SignupPayload) (bool, error) {
	type NewUserRequest struct {
		FirstName string `validate:"min=8"`
		LastName  string `validate:"min=8"`
		UserName  string `validate:"min=8"`
		Email     string `validate:"min=8"`
		Language  string `validate:"min=8"`
		Password  string `validate:"min=8"`
	}

	nur := &NewUserRequest{
		FirstName: signupPayload.FirstName,
		LastName:  signupPayload.LastName,
		UserName:  signupPayload.UserName,
		Email:     signupPayload.Email,
		Language:  signupPayload.Language,
		Password:  signupPayload.Password,
	}
	var validate *validator.Validate

	validate = validator.New()
	err := validate.Struct(nur)

	if err != nil {

		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		if _, ok := err.(*validator.InvalidValidationError); ok {
			log.Println(err)
			return true, errors.New("asd")
		}

		for _, err := range err.(validator.ValidationErrors) {
			log.Printf("%s %s %s %s - %s\n", err.Namespace(), err.Type(), err.Tag(), err.Param(), err.Value())
		}

		// from here you can create your own error messages in whatever language you wish
		return true, errors.New("asd")
	}
	return true, errors.New("asd")
}