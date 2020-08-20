package account

import (
	"github.com/cotrutatiberiu/project-go/service"
)

type Controller struct {
	repo service.Account
}

func New(acct service.Account) *Controller {
	return &Controller{
		repo: acct,
	}
}
