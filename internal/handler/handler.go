package handler

import (
	"note-system/internal/service"
	"note-system/pkg/logging"
	"time"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const (
	responseTimeout = time.Second * 3
)

type Handler struct {
	logger  *logging.Logger
	service *service.Service
}

func NewHandler(logger *logging.Logger, service *service.Service) *Handler {

	return &Handler{
		logger:  logger,
		service: service,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	h.logger.Info("init routes")
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-in", h.signIn)
		auth.POST("/sign-up", h.signUp)
	}

	api := router.Group("/api", h.accountIdentity)
	{
		api.GET("/heartbeat", h.heartbeat)

		note := api.Group("/note")
		{
			note.GET("/", h.getAllNotes)
			note.GET("/:id", h.getNoteById)
			note.POST("/", h.createNote)
			note.DELETE("/:id", h.deleteNote)
			note.PATCH("/:id", h.updateNote)
		}
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
