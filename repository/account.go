package repository

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/cotrutatiberiu/project-go/models"
)

type Account interface {
	// Find(ctx context.Context, id uint64) (*models.Account, error)
	// FindAll(ctx context.Context) ([]models.Account, error)
	// FindByUsername(ctx context.Context, username string) (*models.Account, error)
	// FindByEmail(ctx context.Context, email string) (*models.Account, error)
	Create(conn *pgxpool.Pool, ctx context.Context, acct *models.Account) error
	// Update(ctx context.Context, acct *models.Account) error
	// Delete(ctx context.Context, id uint64) error
}
