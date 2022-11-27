package handler

import (
	"context"
	"net/http"
	"note-system/internal/domain"

	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(c *gin.Context) {
	h.logger.Debugln("signing up account")
	ctx, cancel := context.WithTimeout(c, responseTimeout)
	defer cancel()

	accountDTO := domain.CreateAccountDTO{}

	if err := c.BindJSON(&accountDTO); err != nil {
		ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	accountId, err := h.service.Authorization.SignUp(ctx, accountDTO)
	if err != nil {
		ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	IdResponse(c, http.StatusCreated, accountId)
}
func (h *Handler) signIn(c *gin.Context) {
	h.logger.Debugln("signing in account")
	ctx, cancel := context.WithTimeout(c, responseTimeout)
	defer cancel()

	accountDTO := domain.LoginAccountDTO{}

	if err := c.BindJSON(&accountDTO); err != nil {
		ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.service.Authorization.SignIn(ctx, accountDTO)
	if err != nil {
		ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	h.logger.Debugf("token: %s", token)

	c.JSON(200, map[string]interface{}{"token": token})
}
