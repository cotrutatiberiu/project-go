package account

import (
	"github.com/cotrutatiberiu/project-go/service"
)

type Controller struct {
	serviceAccount service.Account
}

func New(acct service.Account) *Controller {
	return &Controller{
		serviceAccount: acct,
	}
}