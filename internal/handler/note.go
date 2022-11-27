package handler

import (
	"context"
	"net/http"
	"note-system/internal/domain"

	"github.com/gin-gonic/gin"
)

// @Summary Get note by id
// @Tags Note
// @Param id path int true "Note ID"
// @Success 200
// @Failure 400
// @Router /api/note/{id} [get]
func (h *Handler) getById(c *gin.Context) {
	id := c.GetInt("id")
	h.logger.Infof("getting note id:%d", id)
	ctx, cancel := context.WithTimeout(c, responseTimeout)
	defer cancel()

	note, err := h.service.Note.GetById(ctx, 1)
	if err != nil {
		ErrorResponse(c, 403, err.Error())
	}

	c.JSON(204, map[interface{}]int{"value": note})
}

// @Summary Get all user's notes
// @Tags Note
// @Success 200
// @Failure 400
// @Router /api/note/ [get]
func (h *Handler) getAll(c *gin.Context) {

}

// @Summary Create note
// @Tags Note
// @Success 201
// @Failure 400
// @Router /api/note/ [post]
func (h *Handler) create(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, responseTimeout)
	defer cancel()

	accountId, err := h.getAccountId(c)
	if err != nil {
		ErrorResponse(c, http.StatusInternalServerError, "failed to get account id")
		return
	}

	dto := domain.CreateNoteDTO{AccountId: accountId}

	if err := c.BindJSON(&dto); err != nil {
		ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	noteId, err := h.service.Note.Create(ctx, dto)
	if err != nil {
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
func (h *Handler) update(c *gin.Context) {

}

// @Summary Delete note
// @Tags Note
// @Param id path int true "Note ID"
// @Success 200
// @Failure 400
// @Router /api/note/{id} [delete]
func (h *Handler) delete(c *gin.Context) {

}
