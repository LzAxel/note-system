package psql

import (
	"context"
	"database/sql"
	"fmt"
	"note-system/internal/domain"
	"note-system/pkg/logging"
	"time"

	"github.com/Masterminds/squirrel"
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

func (p *NotePostgres) GetById(ctx context.Context, noteDTO domain.GetDeleteNoteDTO) (domain.Note, error) {
	var note domain.Note

	p.logger.Debugf("getting note: id=%d, accountId=%d", noteDTO.Id, noteDTO.AccountId)

	query := fmt.Sprintf("SELECT * FROM %s WHERE (id=$1 AND account_id=$2) OR (id=$3 AND is_public=True)", noteTable)
	if err := p.db.Get(&note, query, noteDTO.Id, noteDTO.AccountId, noteDTO.Id); err != nil {
		return note, err
	}

	return note, nil
}

func (p *NotePostgres) GetAll(ctx context.Context, accountId int) ([]domain.Note, error) {
	var notes = make([]domain.Note, 0)

	query := fmt.Sprintf("SELECT * FROM %s WHERE account_id=$1", noteTable)
	if err := p.db.Select(&notes, query, accountId); err != nil {
		return notes, err
	}

	return notes, nil
}

func (p *NotePostgres) Create(ctx context.Context, note domain.Note) (int, error) {
	var noteId int

	query := fmt.Sprintf("INSERT INTO %s(name, text, tag, url, is_public, account_id) VALUES ($1, $2, $3, $4, $5, $6) RETURNING ID", noteTable)
	if err := p.db.Get(&noteId, query, note.Name, note.Text, note.Tag, note.Url, note.IsPublic, note.AccountId); err != nil {
		return 0, err
	}

	return noteId, nil
}

func (p *NotePostgres) Delete(ctx context.Context, noteDTO domain.GetDeleteNoteDTO) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1 and account_id=$2", noteTable)
	result, err := p.db.Exec(query, noteDTO.Id, noteDTO.AccountId)
	if err != nil {
		return err
	}
	if c, _ := result.RowsAffected(); c == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (p *NotePostgres) Update(ctx context.Context, noteDTO domain.UpdateNoteDTO) (int, error) {
	query := squirrel.Update(noteTable).PlaceholderFormat(squirrel.Dollar)

	if noteDTO.Name != nil {
		query = query.Set("name", noteDTO.Name)
	}
	if noteDTO.Text != nil {
		query = query.Set("text", noteDTO.Text)
	}
	if noteDTO.Tag != nil {
		query = query.Set("tag", noteDTO.Tag)
	}
	if noteDTO.IsPublic != nil {
		query = query.Set("is_public", noteDTO.IsPublic)
	}
	query.Set("updated_at", time.Now())
	queryStr, queryArgs := query.Where("id = ? and account_id = ?", noteDTO.Id, noteDTO.AccountId).MustSql()

	p.logger.Debugf("update note with query: %s | %v", queryStr, queryArgs)

	result, err := p.db.Exec(queryStr, queryArgs...)
	if err != nil {
		return 0, err
	}
	if c, _ := result.RowsAffected(); c == 0 {
		return 0, sql.ErrNoRows
	}
	p.logger.Debug(result.RowsAffected())

	return noteDTO.Id, nil
}
