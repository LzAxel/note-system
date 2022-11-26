package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authHeader = "Authorization"
)

func (h *Handler) accountIdentity(ctx *gin.Context) {
	header := ctx.GetHeader("Authorization")
	if header == "" {
		newErrorResponse(ctx, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		newErrorResponse(ctx, http.StatusUnauthorized, "invalid auth header")
		return
	}

	headerToken := headerParts[1]
	token, err := h.service.JWTManager.Parse(headerToken)
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	claims, err := h.service.JWTManager.Claims(token)
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, "invalid token claims")
		return
	}

	ctx.Set("accountId", claims["sub"].(string))
}
