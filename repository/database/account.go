package database

import (
	"context"
	"errors"

	"github.com/cotrutatiberiu/project-go/models"
	"github.com/cotrutatiberiu/project-go/repository"
	"github.com/jackc/pgx/v4"
)

const (
	sqlFindByID       = "SELECT account_id,email,password FROM accounts WHERE account_id=$1"
	sqlFindByUsername = "SELECT account_id,email,password FROM accounts WHERE username=$1"
	sqlFindByEmail    = "SELECT account_id,email,password FROM accounts WHERE email=$1"
)

type account struct {
	db *pgx.Conn
}

// Account return new account repository
func Account(conn *pgx.Conn) repository.Account {
	return &account{db: conn}
}

// Find account by id
func (r *account) Find(ctx context.Context, id uint64) (*models.Account, error) {
	row := r.db.QueryRow(ctx, sqlFindByID, id)
	acct := models.Account{}
	err := row.Scan(&acct.AccountID, &acct.Email, &acct.Password)
	if err != nil {
		return nil, err
	}
	return &acct, nil
}

// FindAll return a  list of accounts
func (r *account) FindAll(ctx context.Context) ([]models.Account, error) {
	return nil, errors.New("Not implemented")
}

// FindByUsername like Find account but using username as search criteria
func (r *account) FindByUsername(ctx context.Context, username string) (*models.Account, error) {
	row := r.db.QueryRow(ctx, sqlFindByUsername, username)
	acct := models.Account{}
	err := row.Scan(&acct.AccountID, &acct.Email, &acct.Password)
	if err != nil {
		return nil, err
	}
	return &acct, nil
}

// FindByEmail like Find account but using email as search criteria
func (r *account) FindByEmail(ctx context.Context, email string) (*models.Account, error) {
	row := r.db.QueryRow(ctx, sqlFindByEmail, email)
	acct := models.Account{}
	err := row.Scan(&acct.AccountID, &acct.Email, &acct.Password)
	if err != nil {
		return nil, err
	}
	return &acct, nil
}

// Create account record
func (r *account) Create(ctx context.Context, acct *models.Account) error {
	//result := db.Create(pool, `INSERT INTO
	//		accounts(first_name, last_name, user_name, email, language, password, created, updated)
	//		VALUES($1, $2, $3, $4, $5, $6, $7, $8);`)
	return errors.New("Not implemented")
}

// Update account by ID
func (r *account) Update(ctx context.Context, acct *models.Account) error {
	return errors.New("Not implemented")
}

// Delete account by ID
func (r *account) Delete(ctx context.Context, id uint64) error {
	return errors.New("Not implemented")
}
