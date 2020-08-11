package repository

import (
	"context"

	"github.com/cotrutatiberiu/project-go/models"
)

type Account interface {
	Find(ctx context.Context, id uint64) (*models.Account, error)
	FindAll(ctx context.Context) ([]models.Account, error)
	ByEmail(ctx context.Context, email string) (*models.Account, error)
	Create(ctx context.Context, acct *models.Account) error
	Update(ctx context.Context, acct *models.Account) error
	Delete(ctx context.Context, id uint64) error
}
