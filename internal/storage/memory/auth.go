package memory

import (
	"context"
	"errors"
	"note-system/internal/domain"
	"note-system/pkg/logging"
	"sync"
)

var counter int = 0

type AuthMemory struct {
	m      sync.Map
	logger *logging.Logger
}

func NewAuthMemory(logger *logging.Logger) *AuthMemory {
	return &AuthMemory{logger: logger}
}

func (p *AuthMemory) SignUp(ctx context.Context, account domain.Account) (int, error) {
	p.logger.Debugln("storage signing up accont")
	account.Id = counter
	p.m.Store(account.Username, account)
	counter++

	return account.Id, nil
}

func (p *AuthMemory) SignIn(ctx context.Context, accountDTO domain.LoginAccountDTO) (domain.Account, error) {
	p.logger.Debugf("getting from database: %s", accountDTO.Username)
	var account domain.Account
	loadedAccount, ok := p.m.Load(accountDTO.Username)
	if !ok {
		return account, errors.New("account not found")
	}
	account, ok = loadedAccount.(domain.Account)
	if !ok {
		return account, errors.New("account not found")
	}

	return account, nil
}
