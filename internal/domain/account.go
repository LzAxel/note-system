package domain

import (
	"errors"
	"time"
)

var (
	ErrNotFound    = errors.New("not found")
	ErrFailedToGet = errors.New("failed to get")
	ErrNotTheOwner = errors.New("not the owner of note")
)

type Account struct {
	Id           int       `db:"id"`
	Username     string    `db:"username"`
	PasswordHash string    `db:"password_hash"`
	HashSalt     string    `db:"hash_salt"`
	CreatedAt    time.Time `db:"created_at"`
}

type CreateAccountDTO struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginAccountDTO struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
