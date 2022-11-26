package handler

import (
	"github.com/gin-gonic/gin"
)

type errorResponse struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, code int, message string) {
	c.AbortWithStatusJSON(code, errorResponse{Message: message})
}

type idResponse struct {
	Id int `json:"id"`
}

func newIdResponse(c *gin.Context, code int, id int) {
	c.JSON(code, idResponse{Id: id})
}
