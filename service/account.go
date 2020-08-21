package service

import (
	"context"

	"github.com/cotrutatiberiu/project-go/models"
)

type Account interface {
	Create(ctx context.Context, account *models.Account) error
	SetPassword(account *models.Account, password string) error
}
