package account

import (
	"context"

	"github.com/cotrutatiberiu/project-go/models"
	"github.com/cotrutatiberiu/project-go/repository"
	"github.com/cotrutatiberiu/project-go/service"
)

type svc struct {
	account repository.Account
}

func New(repo repository.Account) service.Account {
	return &svc{account: repo}
}

func (s *svc) Create(ctx context.Context, account *models.Account) error {
	err := account.Validate()
	if err != nil {
		return err
	}
	return s.account.Create(ctx, account)
}

func (s *svc) SetPassword(account *models.Account, password string) error {
	return nil
}
