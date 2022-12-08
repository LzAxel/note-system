package memory

import (
	"context"
	"note-system/internal/domain"
	"note-system/pkg/logging"
	"sync"
	"time"
)

var noteCounter int = 0

type NoteMemory struct {
	m      sync.Map
	logger *logging.Logger
}

func NewNoteMemory(logger *logging.Logger) *NoteMemory {
	return &NoteMemory{logger: logger}
}

func (m *NoteMemory) GetById(ctx context.Context, noteDTO domain.GetDeleteNoteDTO) (domain.Note, error) {
	var note domain.Note

	m.logger.Debugf("getting note: id=%v, accountId=%v", noteDTO.Id, noteDTO.AccountId)

	note, err := m.loadNoteById(noteDTO.Id)
	if err != nil {
		return note, err
	}
	m.logger.Debugf("isPublic: %v, accountId: %v", note.IsPublic, note.AccountId)
	if note.AccountId != noteDTO.AccountId && note.IsPublic != true {
		return note, domain.ErrNotTheOwner
	}

	return note, nil
}

func (m *NoteMemory) GetAll(ctx context.Context, accountId int) ([]domain.Note, error) {
	var notes = make([]domain.Note, 0)

	m.m.Range(func(key, value any) bool {
		n, ok := value.(domain.Note)
		if !ok {
			return true
		}

		if n.AccountId == accountId {
			notes = append(notes, n)
		}

		return true
	})

	return notes, nil
}

func (m *NoteMemory) Create(ctx context.Context, note domain.Note) (int, error) {
	note.Id = noteCounter
	note.UpdatedAt = time.Now()
	note.CreatedAt = time.Now()
	m.m.Store(note.Id, note)
	noteCounter++

	return note.Id, nil
}

func (m *NoteMemory) Delete(ctx context.Context, noteDTO domain.GetDeleteNoteDTO) error {
	var note domain.Note

	note, err := m.loadNoteById(noteDTO.Id)
	if err != nil {
		return err
	}
	if note.AccountId != noteDTO.AccountId {
		return domain.ErrNotTheOwner
	}
	m.m.Delete(note.Id)

	return nil
}

func (m *NoteMemory) Update(ctx context.Context, noteDTO domain.UpdateNoteDTO) (domain.Note, error) {
	var note domain.Note

	note, err := m.loadNoteById(noteDTO.Id)
	if err != nil {
		return note, err
	}
	if note.AccountId != noteDTO.AccountId {
		return note, domain.ErrNotTheOwner
	}

	if noteDTO.Name != nil {
		note.Name = *noteDTO.Name
	}
	if noteDTO.Text != nil {
		note.Text = *noteDTO.Text
	}
	if noteDTO.Tag != nil {
		note.Tag = *noteDTO.Tag
	}
	if noteDTO.IsPublic != nil {
		note.IsPublic = *noteDTO.IsPublic
	}
	note.UpdatedAt = time.Now()
	m.m.Store(note.Id, note)

	return note, nil
}

func (m *NoteMemory) loadNoteById(id int) (domain.Note, error) {
	var note domain.Note

	val, ok := m.m.Load(id)
	if !ok {
		return note, domain.ErrNotFound
	}
	note, ok = val.(domain.Note)
	if !ok {
		return note, domain.ErrFailedToGet
	}

	return note, nil
}
