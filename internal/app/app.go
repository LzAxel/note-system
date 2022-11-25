package app

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"note-system/internal/config"
	"note-system/pkg/logging"
	"note-system/pkg/metric"
	"time"

	_ "note-system/docs"

	"github.com/julienschmidt/httprouter"
	httpSwagger "github.com/swaggo/http-swagger"
)

type App struct {
	cfg        *config.Config
	logger     *logging.Logger
	router     *httprouter.Router
	httpServer *http.Server
}

func (a *App) Run() {
	a.startHTTP()
}

func NewApp(config *config.Config, logger *logging.Logger) (App, error) {
	logger.Info("router init")
	router := httprouter.New()

	logger.Info("swagger docs init")
	router.Handler(http.MethodGet, "/swagger",
		http.RedirectHandler("/swagger/index.html", http.StatusPermanentRedirect))
	router.Handler(http.MethodGet, "/swagger/*any", httpSwagger.WrapHandler)

	logger.Info("heartbeat init")
	metricHandler := metric.Handler{}
	metricHandler.Register(router)

	app := App{
		cfg:    config,
		logger: logger,
		router: router,
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
		Handler:      a.router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
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
