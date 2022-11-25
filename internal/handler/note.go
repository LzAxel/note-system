package handler

import (
	"github.com/gin-gonic/gin"
)

// @Summary Get note by id
// @Tags Note
// @Param id path int true "Note ID"
// @Success 200
// @Failure 400
// @Router /api/note/{id} [get]
func (h *Handler) getById(c *gin.Context) {

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
// @Success 200
// @Failure 400
// @Router /api/note/ [post]
func (h *Handler) create(c *gin.Context) {

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
