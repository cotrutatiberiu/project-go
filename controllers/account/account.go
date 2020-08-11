package account

import (
	"github.com/cotrutatiberiu/project-go/repository"
)

type Controller struct {
	repo repository.Account
}

func New(acct repository.Account) *Controller {
	return &Controller{
		repo: acct,
	}
}
