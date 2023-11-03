package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	// Swagger docs.
	_ "llm_plateform_resourcemanagement/docs"
	"llm_plateform_resourcemanagement/internal/usecase"
	"llm_plateform_resourcemanagement/pkg/logger"
)

// @title swagger test
// @version 1.0
// @description swagger test example
// @schemes http https
// @BasePath /v1
func NewRouter(handler *gin.Engine, l logger.Interface, t usecase.Test) {
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	swaggerHandler := ginSwagger.WrapHandler(swaggerFiles.Handler)
	handler.GET("/swagger/*any", swaggerHandler)

	handler.GET("/healthz", func(c *gin.Context) { c.Status(http.StatusOK) })

	handler.GET("/metrics", gin.WrapH(promhttp.Handler()))

	h := handler.Group("/v1")
	{
		newTestRoutes(h, t, l)
	}
}
