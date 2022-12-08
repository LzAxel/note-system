package handler

import (
	"context"
	"net/http"
	"note-system/internal/domain"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Get note by id
// @Tags Note
// @Param id path int true "Note ID"
// @Success 200
// @Failure 400
// @Router /api/note/{id} [get]
func (h *Handler) getNoteById(c *gin.Context) {
	var note domain.Note

	accountId, err := h.getAccountId(c)
	if err != nil {
		h.logger.Errorf(err.Error())
		ErrorResponse(c, 403, err.Error())
		return
	}

	id := c.Param("id")
	h.logger.Infof("getting note id:%d", id)
	idInt, err := strconv.Atoi(id)
	if err != nil {
		h.logger.Errorf(err.Error())
		ErrorResponse(c, http.StatusBadRequest, "invalid note id")
		return
	}
	ctx, cancel := context.WithTimeout(c, responseTimeout)
	defer cancel()

	noteDTO := domain.GetDeleteNoteDTO{Id: idInt, AccountId: accountId}

	note, err = h.service.Note.GetById(ctx, noteDTO)
	if err != nil {
		h.logger.Errorf(err.Error())
		ErrorResponse(c, 403, err.Error())
		return
	}

	c.JSON(200, note)
}

// @Summary Get all user's notes
// @Tags Note
// @Success 200
// @Failure 400
// @Router /api/note/ [get]
func (h *Handler) getAllNotes(c *gin.Context) {
	var notes = make([]domain.Note, 0)

	accountId, err := h.getAccountId(c)
	if err != nil {
		h.logger.Errorf(err.Error())
		ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	h.logger.Infof("getting note:%d", accountId)
	ctx, cancel := context.WithTimeout(c, responseTimeout)
	defer cancel()

	notes, err = h.service.Note.GetAll(ctx, accountId)
	if err != nil {
		h.logger.Errorf(err.Error())
		ErrorResponse(c, 403, err.Error())
		return
	}

	c.JSON(200, notes)
}

// @Summary Create note
// @Tags Note
// @Success 201
// @Failure 400
// @Router /api/note/ [post]
func (h *Handler) createNote(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, responseTimeout)
	defer cancel()

	accountId, err := h.getAccountId(c)
	if err != nil {
		h.logger.Errorf(err.Error())
		ErrorResponse(c, http.StatusInternalServerError, "failed to get account id")
		return
	}

	dto := domain.CreateNoteDTO{AccountId: accountId}

	if err := c.BindJSON(&dto); err != nil {
		h.logger.Errorf(err.Error())
		ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	noteId, err := h.service.Note.Create(ctx, dto)
	if err != nil {
		h.logger.Errorf(err.Error())
		ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	IdResponse(c, 201, noteId)
}

// @Summary Update note
// @Tags Note
// @Success 200
// @Failure 400
// @Router /api/note/{id} [post]
func (h *Handler) updateNote(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, responseTimeout)
	defer cancel()

	accountId, err := h.getAccountId(c)
	if err != nil {
		h.logger.Errorf(err.Error())
		ErrorResponse(c, http.StatusInternalServerError, "failed to get account id")
		return
	}
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		h.logger.Errorf(err.Error())
		ErrorResponse(c, http.StatusBadRequest, "invalid note id")
		return
	}
	dto := domain.UpdateNoteDTO{AccountId: accountId, Id: idInt}

	if err := c.BindJSON(&dto); err != nil {
		h.logger.Errorf(err.Error())
		ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	note, err := h.service.Note.Update(ctx, dto)
	if err != nil {
		h.logger.Errorf(err.Error())
		ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, note)
}

// @Summary Delete note
// @Tags Note
// @Param id path int true "Note ID"
// @Success 200
// @Failure 400
// @Router /api/note/{id} [delete]
func (h *Handler) deleteNote(c *gin.Context) {
	accountId, err := h.getAccountId(c)
	if err != nil {
		h.logger.Errorf(err.Error())
		ErrorResponse(c, 403, err.Error())
		return
	}

	id := c.Param("id")
	h.logger.Infof("deleting note id:%v", id)
	idInt, err := strconv.Atoi(id)
	if err != nil {
		h.logger.Errorf(err.Error())
		ErrorResponse(c, http.StatusBadRequest, "invalid note id")
		return
	}
	ctx, cancel := context.WithTimeout(c, responseTimeout)
	defer cancel()

	noteDTO := domain.GetDeleteNoteDTO{Id: idInt, AccountId: accountId}

	if err = h.service.Note.Delete(ctx, noteDTO); err != nil {
		h.logger.Errorf(err.Error())
		ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusNoContent)
}
