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

func (s *NoteService) GetById(ctx context.Context, noteDTO domain.GetDeleteNoteDTO) (domain.Note, error) {
	return s.storage.GetById(ctx, noteDTO)
}
func (s *NoteService) GetAll(ctx context.Context, accountId int) ([]domain.Note, error) {
	return s.storage.GetAll(ctx, accountId)
}

func (s *NoteService) Create(ctx context.Context, noteDTO domain.CreateNoteDTO) (int, error) {
	note := domain.Note{
		Name:      noteDTO.Name,
		Text:      noteDTO.Text,
		Tag:       noteDTO.Tag,
		Url:       "testnoteurl222",
		IsPublic:  noteDTO.IsPublic,
		AccountId: noteDTO.AccountId,
	}

	s.logger.Info(note)
	return s.storage.Create(ctx, note)
}

func (s *NoteService) Delete(ctx context.Context, noteDTO domain.GetDeleteNoteDTO) error {
	return s.storage.Delete(ctx, noteDTO)
}

func (p *NoteService) Update(ctx context.Context, noteDTO domain.UpdateNoteDTO) (int, error) {
	return p.storage.Update(ctx, noteDTO)
}
