package account

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/cotrutatiberiu/project-go/models"
	"github.com/cotrutatiberiu/project-go/repository"
)

const (
	sqlCreateAccount = `INSERT INTO
		accounts(first_name, last_name, user_name, email, language, password, created, updated)
		VALUES($1, $2, $3, $4, $5, $6, $7, $8);`
)

type account struct{}

// Account return new account repository
func AccountRepository() repository.Account {
	return &account{}
}

type dbobj struct {
	first_name  string
	last_name   string
	user_name   string
	email       string
	language    string
	password    string
	active      bool
	admin_level int64
	created     int64
}

func (a *account) Create(conn *pgxpool.Pool, ctx context.Context, account *models.Account) error {

	tag, err := conn.Exec(ctx,
		sqlCreateAccount,
		account.FirstName,
		account.LastName,
		account.UserName,
		account.Email,
		account.Language,
		account.Password,
		account.Created,
		account.Updated)
	fmt.Println(err)
	fmt.Println(tag)

	return err
}

func Read(pool *pgxpool.Pool, ctx context.Context, username string, password string) {

}

// func Update() {

// }

// func Delete() {

// }
