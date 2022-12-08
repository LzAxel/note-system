package memory

import (
	"context"
	"note-system/internal/domain"
	"note-system/pkg/logging"
	"sync"
	"time"
)

var authCounter int = 0

type AuthMemory struct {
	m      sync.Map
	logger *logging.Logger
}

func NewAuthMemory(logger *logging.Logger) *AuthMemory {
	return &AuthMemory{logger: logger}
}

func (p *AuthMemory) SignUp(ctx context.Context, account domain.Account) (int, error) {
	p.logger.Debugln("storage signing up accont")
	account.Id = authCounter
	account.CreatedAt = time.Now()
	p.m.Store(account.Username, account)
	authCounter++

	return account.Id, nil
}

func (p *AuthMemory) SignIn(ctx context.Context, accountDTO domain.LoginAccountDTO) (domain.Account, error) {
	p.logger.Debugf("getting from database: %s", accountDTO.Username)
	var account domain.Account
	loadedAccount, ok := p.m.Load(accountDTO.Username)
	if !ok {
		return account, domain.ErrNotFound
	}
	account, ok = loadedAccount.(domain.Account)
	if !ok {
		return account, domain.ErrFailedToGet
	}

	return account, nil
}
