package psql

import (
	"context"
	"fmt"
	"note-system/internal/domain"
	"note-system/pkg/logging"

	"github.com/jmoiron/sqlx"
)

const (
	accountTable = "account"
)

type AuthPostgres struct {
	db     *sqlx.DB
	logger *logging.Logger
}

func NewAuthPostgres(db *sqlx.DB, logger *logging.Logger) *AuthPostgres {
	return &AuthPostgres{db: db, logger: logger}
}

func (p *AuthPostgres) SignUp(ctx context.Context, account domain.Account) (int, error) {
	p.logger.Debugln("storage signing up accont")

	var accountId int

	query := fmt.Sprintf("INSERT INTO %s(username, password_hash, hash_salt) VALUES ($1, $2, $3) RETURNING ID",
		accountTable)
	if err := p.db.Get(&accountId, query, account.Username, account.PasswordHash,
		account.HashSalt); err != nil {

		return 0, err
	}

	return accountId, nil
}

func (p *AuthPostgres) SignIn(ctx context.Context, accountDTO domain.LoginAccountDTO) (domain.Account, error) {
	p.logger.Debugf("getting from database: %s", accountDTO.Username)
	var account domain.Account

	query := fmt.Sprintf("SELECT * FROM %s WHERE username=$1",
		accountTable)
	if err := p.db.Get(&account, query, accountDTO.Username); err != nil {
		return account, err
	}

	return account, nil
}
