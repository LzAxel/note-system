package handler

import (
	"note-system/pkg/logging"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	logger *logging.Logger
}

func NewHandler(logger *logging.Logger) *Handler {
	return &Handler{
		logger: logger,
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

	api := router.Group("/api")
	{
		api.GET("/heartbeat", h.Heartbeat)

		note := api.Group("/note")
		{
			note.GET("/", h.getAll)
			note.GET("/:id", h.getById)
			note.POST("/", h.create)
			note.DELETE("/:id", h.delete)
			note.PATCH("/:id", h.update)
		}
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
