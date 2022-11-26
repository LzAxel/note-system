package service

import (
	"context"
	"note-system/internal/domain"
	"note-system/internal/storage"
	"note-system/pkg/jwt"
	"note-system/pkg/logging"
)

type Authorization interface {
	SignUp(ctx context.Context, accountDTO domain.CreateAccountDTO) (int, error)
	SignIn(ctx context.Context, accountDTO domain.LoginAccountDTO) (string, error)
}

type Note interface {
	GetById(ctx context.Context, id int) (int, error)
	Create(ctx context.Context, noteDTO domain.CreateNoteDTO) (int, error)
}

type Service struct {
	Authorization
	Note

	storage    *storage.Storage
	logger     *logging.Logger
	JWTManager *jwt.JWTManager
}

func NewService(logger *logging.Logger, storage *storage.Storage, manager *jwt.JWTManager) *Service {
	return &Service{
		storage:       storage,
		logger:        logger,
		JWTManager:    manager,
		Note:          NewNoteService(storage, logger),
		Authorization: NewAuthService(storage, logger, manager),
	}
}
