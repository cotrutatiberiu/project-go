package signup

import (
	"github.com/cotrutatiberiu/project-go/service"
)

type Controller struct {
	accountService service.Account
	validationService service.Validation
}

func New(accountService service.Account,validationService service.Validation) *Controller {
	return &Controller{accountService,validationService}
}
