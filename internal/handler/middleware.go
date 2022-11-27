package handler

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authHeader = "Authorization"
)

func (h *Handler) accountIdentity(ctx *gin.Context) {
	header := ctx.GetHeader("Authorization")
	if header == "" {
		ErrorResponse(ctx, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		ErrorResponse(ctx, http.StatusUnauthorized, "invalid auth header")
		return
	}

	headerToken := headerParts[1]
	token, err := h.service.JWTManager.Parse(headerToken)
	if err != nil {
		ErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	claims, err := h.service.JWTManager.Claims(token)
	if err != nil {
		ErrorResponse(ctx, http.StatusUnauthorized, "invalid token claims")
		return
	}

	ctx.Set("accountId", claims["sub"].(string))
}

func (h *Handler) getAccountId(ctx *gin.Context) (int, error) {
	accountId := ctx.Value("accountId").(string)
	if accountId == "" {
		return 0, errors.New("failed to get account id")
	}
	accountIdInt, err := strconv.Atoi(accountId)
	if err != nil {
		return 0, errors.New("failed to get account id")
	}

	return accountIdInt, nil
}
