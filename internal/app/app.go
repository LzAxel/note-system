package app

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"note-system/internal/config"
	"note-system/internal/handler"
	"note-system/pkg/logging"
	"time"

	_ "note-system/docs"

	"github.com/gin-gonic/gin"
)

type App struct {
	cfg        *config.Config
	logger     *logging.Logger
	router     *gin.Engine
	httpServer *http.Server
}

func (a *App) Run() {
	a.startHTTP()
}

func NewApp(config *config.Config, logger *logging.Logger) (App, error) {
	router := handler.NewHandler(logger)

	app := App{
		cfg:    config,
		logger: logger,
		router: router.InitRoutes(),
	}

	return app, nil
}

func (a *App) startHTTP() {
	a.logger.Info("start HTTP")

	var listener net.Listener

	a.logger.Infof("bind to host: %s, port: %s", a.cfg.Listen.BindIP, a.cfg.Listen.Port)

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", a.cfg.Listen.BindIP, a.cfg.Listen.Port))
	if err != nil {
		a.logger.Fatal(err)
	}

	a.httpServer = &http.Server{
		Handler:        a.router,
		MaxHeaderBytes: 1 << 20,
		WriteTimeout:   15 * time.Second,
		ReadTimeout:    15 * time.Second,
	}

	a.logger.Info("app started")

	if err := a.httpServer.Serve(listener); err != nil {
		switch {
		case errors.Is(err, http.ErrServerClosed):
			a.logger.Warn("server shutdown")
		default:
			a.logger.Fatal(err)
		}
	}
	err = a.httpServer.Shutdown(context.Background())
	if err != nil {
		a.logger.Fatal(err)
	}
}
