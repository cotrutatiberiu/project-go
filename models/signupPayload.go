package models

import "time"

type SignupPayload struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	UserName  string `json:"userName"`
	Email     string `json:"email"`
	Language  string `json:"language"`
	Password  string `json:"password"`
	// Created   int64  `json:"created"`
}

func (payload SignupPayload) ToModel() *Account {
	return &Account{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		UserName:  payload.UserName,
		Email:     payload.Email,
		Language:  payload.Language,
		Password:  payload.Password,
		Created:   time.Now(),
		Updated:   time.Now(),
	}
}
