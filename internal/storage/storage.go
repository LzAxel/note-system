package storage

import (
	"context"
	"note-system/internal/domain"
	"note-system/internal/storage/psql"
	"note-system/pkg/logging"

	"github.com/jmoiron/sqlx"
)

type Auth interface {
	SignUp(ctx context.Context, account domain.Account) (int, error)
	SignIn(ctx context.Context, accountDTO domain.LoginAccountDTO) (domain.Account, error)
}

type Note interface {
	GetById(ctx context.Context, noteId int) (int, error)
	Create(ctx context.Context, note domain.Note) (int, error)
}

type Storage struct {
	Auth
	Note

	logger *logging.Logger
}

func NewStorage(logger *logging.Logger, db *sqlx.DB) *Storage {
	return &Storage{
		logger: logger,
		Note:   psql.NewNotePostgres(db, logger),
		Auth:   psql.NewAuthPostgres(db, logger),
	}
}
