package session

import (
	"context"
	"errors"
	"time"

	"github.com/cotrutatiberiu/project-go/repository"
	"github.com/cotrutatiberiu/project-go/service"
)

type svc struct {
	privateKey     string        // JWT private key
	expireDuration time.Duration // Token expire time
	account        repository.Account
}

func New(repo repository.Account) service.Session {
	return &svc{
		expireDuration: 3600 * time.Second,
		account:        repo,
	}
}

// Login service
func (s *svc) Login(ctx context.Context, username, password string) (string, error) {

	acct, err := s.account.FindByUsername(ctx, username)
	if err != nil {
		return "", err
	}

	// TODO check hasked password
	if acct.Password != password {
		return "", errors.New("password don't match")
	}

	// TODO generate session token
	token := "K4HGd4h3JhH1KX4pzACT5om1yn0M1c"

	return token, nil
}
