package account

import (
	"context"
	"errors"

	"github.com/cotrutatiberiu/project-go/models"
	"github.com/cotrutatiberiu/project-go/repository"
	"github.com/cotrutatiberiu/project-go/repository/database"
	"github.com/cotrutatiberiu/project-go/service"
)

type accountService struct {
	accountRepository repository.Account
	db                *database.Database
}

func New(repo repository.Account, db *database.Database) service.Account {
	return &accountService{repo, db}
}

func (service *accountService) Create(ctx context.Context, account *models.Account) error {
	// err := account.Validate()
	// if err != nil {
	// 	return err
	// }

	conn := service.db.CreateConnection()

	service.accountRepository.Create(conn, ctx, account)

	conn.Close()

	return errors.New("asd")
}
