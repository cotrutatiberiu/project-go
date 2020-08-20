package service

import "context"

// Session service definition
type Session interface {
	Login(ctx context.Context, username, password string) (string, error)
}
