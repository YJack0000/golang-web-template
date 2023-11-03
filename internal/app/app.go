package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"

	"llm_plateform_resourcemanagement/config"
	v1 "llm_plateform_resourcemanagement/internal/controller/http/v1"
	"llm_plateform_resourcemanagement/internal/usecase"
	"llm_plateform_resourcemanagement/internal/usecase/memo"

	"llm_plateform_resourcemanagement/pkg/httpserver"
	"llm_plateform_resourcemanagement/pkg/logger"
)

func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	testUseCase := usecase.New(
		memo.New(),
	)

	handler := gin.New()
	v1.NewRouter(handler, l, testUseCase)
	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err := <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	err := httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}
