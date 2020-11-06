package signup

import (
	"github.com/cotrutatiberiu/project-go/service"
)

type Controller struct {
	accountService service.Account
}

func New(accountService service.Account) *Controller {
	return &Controller{accountService}
}
