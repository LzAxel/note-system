package service

import (
	"context"
	"note-system/internal/domain"
	"note-system/internal/storage"
	"note-system/pkg/logging"
)

type NoteService struct {
	storage storage.Note
	logger  *logging.Logger
}

func NewNoteService(storage storage.Note, logger *logging.Logger) *NoteService {
	return &NoteService{
		storage: storage,
		logger:  logger,
	}
}

func (s *NoteService) GetById(ctx context.Context, noteId int) (int, error) {
	return s.storage.GetById(ctx, noteId)
}

func (s *NoteService) Create(ctx context.Context, noteDTO domain.CreateNoteDTO) (int, error) {
	note := domain.Note{
		Name:      noteDTO.Name,
		Text:      noteDTO.Text,
		Tag:       noteDTO.Tag,
		Url:       "testnoteurl222",
		AccountId: 1,
	}

	s.logger.Info(note)
	return s.storage.Create(ctx, note)
}
