package psql

import (
	"context"
	"fmt"
	"note-system/internal/domain"
	"note-system/pkg/logging"

	"github.com/jmoiron/sqlx"
)

const (
	noteTable = "note"
)

type NotePostgres struct {
	db     *sqlx.DB
	logger *logging.Logger
}

func NewNotePostgres(db *sqlx.DB, logger *logging.Logger) *NotePostgres {
	return &NotePostgres{db: db, logger: logger}
}

func (p *NotePostgres) GetById(ctx context.Context, id int) (int, error) {
	return 3123, nil
}

func (p *NotePostgres) Create(ctx context.Context, note domain.Note) (int, error) {
	var noteId int

	query := fmt.Sprintf("INSERT INTO %s(name, text, tag, url, is_public, account_id) VALUES ($1, $2, $3, $4, $5, $6) RETURNING ID", noteTable)
	if err := p.db.Get(&noteId, query, note.Name, note.Text, note.Tag, note.Url, note.IsPublic, note.AccountId); err != nil {
		return 0, err
	}

	return noteId, nil
}
