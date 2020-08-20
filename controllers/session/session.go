package session

import (
	"github.com/cotrutatiberiu/project-go/service"
)

type Controller struct {
	svc  service.Session
	acct service.Account
}

func New(sessionService service.Session, accountService service.Account) *Controller {
	return &Controller{
		svc:  sessionService,
		acct: accountService,
	}
}
